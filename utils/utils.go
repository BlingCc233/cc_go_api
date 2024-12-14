package utils

import (
	"bytes"
	"fmt"
	"github.com/FloatTech/gg"
	"github.com/golang-jwt/jwt"
	"github.com/skip2/go-qrcode"
	"golang.org/x/crypto/bcrypt"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
	"studyGo/models"
	"time"
	"unicode/utf8"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString([]byte("secret"))
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法是否符合预期
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回 secret，建议从配置文件或环境变量加载，而不是硬编码
		return []byte("secret"), nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	// 提取并验证声明
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// 确保 "username" 是字符串
	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username not found or not a string in token claims")
	}

	return username, nil
}

// GenerateQCSImage creates an image based on QCS fields
func GenerateQCSImage(qcs models.QCS) ([]byte, error) {
	// Load background image
	bgPath := filepath.Join(".", "assets", "qcs_bg.JPG")
	im, err := gg.LoadImage(bgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load background image: %w", err)
	}

	dc := gg.NewContextForImage(im)

	// Load font
	fontPath := filepath.Join(".", "assets", "handwrite.ttf") // Ensure you have a suitable font file
	if err := dc.LoadFontFace(fontPath, 50); err != nil {
		return nil, fmt.Errorf("failed to load font: %w", err)
	}
	fontsd, _ := os.ReadFile(fontPath)

	dc.SetRGB(1, 1, 1) // Red color

	// Helper to draw multi-line text
	drawMultilineText := func(dc *gg.Context, text string, x, y, lineHeight int) error {
		lines := strings.Split(text, "\n")
		for i, line := range lines {
			dc.DrawStringAnchored(line, float64(dc.Width())/2, float64(y+i*lineHeight), 0.5, 0.5)
		}
		return nil
	}

	// Define text positions and draw with different font sizes
	lineHeight := 40 // Line height for multi-line text
	fontSizes := map[string]float64{
		"Luck":           72,
		"Description":    36,
		"Howsas":         25,
		"Interpretation": 20,
		"Default":        100,
	}

	positions := []struct {
		Text string
		X, Y int
		Key  string
	}{
		{Text: qcs.Luck, X: 320, Y: 100, Key: "Luck"},
		{Text: qcs.Description, X: 300, Y: 200, Key: "Description"},
		{Text: qcs.Interpretation, X: 50, Y: 400, Key: "Interpretation"},
		{Text: qcs.Howsas, X: 50, Y: 850, Key: "Howsas"},
		{Text: qcs.Luck, X: 320, Y: 760, Key: "Default"},
	}

	for _, pos := range positions {
		fontSize, exists := fontSizes[pos.Key]
		if !exists {
			fontSize = fontSizes["Default"]
		}
		dc.ParseFontFace(fontsd, fontSize)
		if err := drawMultilineText(dc, pos.Text, pos.X, pos.Y, lineHeight); err != nil {
			return nil, err
		}
	}

	// Encode the image to a buffer
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, dc.Image(), nil); err != nil {
		return nil, fmt.Errorf("failed to encode image: %w", err)
	}

	return buffer.Bytes(), nil
}

// GenerateXBImage creates an image based on goodNew fields
func GenerateXBImage(content string) ([]byte, error) {
	// Initialize gg context
	bgPath := filepath.Join(".", "assets", "xb_bg.png")
	im, err := gg.LoadImage(bgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load background image: %w", err)
	}

	dc := gg.NewContextForImage(im)

	// Load font
	fontPath := filepath.Join(".", "assets", "Songti.ttc")
	if err := dc.LoadFontFace(fontPath, 50); err != nil {
		return nil, fmt.Errorf("failed to load font: %w", err)
	}
	fontsd, _ := os.ReadFile(fontPath)

	// Set font color
	dc.SetRGB(1, 0, 0) // Red color
	dc.ParseFontFace(fontsd, 50)

	var lines []string
	for len(content) > 0 {
		// 获取当前字符的长度
		r, size := utf8.DecodeRuneInString(content)
		if size == 0 {
			break // 如果没有更多字符，退出循环
		}

		// 如果当前行长度加上这个字符的长度超过5，则开始新的一行
		if len(lines) > 0 && len(lines[len(lines)-1]) < 14 {
			lines[len(lines)-1] += string(r)
		} else {
			lines = append(lines, string(r))
		}

		// 移动到下一个字符
		content = content[size:]
	}

	// Calculate text block dimensions
	lineHeight := 60.0
	numLines := len(lines)
	y := 270.0
	y = y - float64(numLines)/2*lineHeight

	// Draw each line centered
	for i, line := range lines {
		dc.DrawStringAnchored(line, float64(dc.Width())/2, y+float64(i)*lineHeight, 0.5, 0.5)
	}

	// Encode the final image to a buffer
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, dc.Image(), nil); err != nil {
		return nil, fmt.Errorf("failed to encode image: %w", err)
	}

	return buffer.Bytes(), nil
}

func QRCode(text string, size int) ([]byte, error) {
	qr, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		return nil, err
	}
	qr.DisableBorder = true
	return qr.PNG(size)
}
