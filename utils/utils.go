package utils

import (
	"bytes"
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"studyGo/models"
	"time"
)

import "github.com/golang-jwt/jwt"

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
	bgFile, err := os.Open(bgPath)
	if err != nil {
		return nil, err
	}
	defer bgFile.Close()

	bgImage, err := jpeg.Decode(bgFile)
	if err != nil {
		return nil, err
	}

	// Create a new RGBA image
	overlay := image.NewRGBA(bgImage.Bounds())
	draw.Draw(overlay, bgImage.Bounds(), bgImage, image.Point{}, draw.Src)

	// Load font
	fontPath := filepath.Join(".", "assets", "handwrite.ttf") // Ensure you have a suitable font file
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}

	ttfFont, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	// Initialize freetype context
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(ttfFont)
	c.SetClip(overlay.Bounds())
	c.SetDst(overlay)
	c.SetSrc(image.White)

	// Helper to draw multi-line text
	drawMultilineText := func(c *freetype.Context, text string, x, y, lineHeight int) error {
		lines := strings.Split(text, "\n")
		for i, line := range lines {
			pt := freetype.Pt(x, y+(i*lineHeight))
			_, err := c.DrawString(line, pt)
			if err != nil {
				return err
			}
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
		{Text: qcs.Howsas, X: 50, Y: 800, Key: "Howsas"},
		{Text: qcs.Luck, X: 320, Y: 760, Key: "Default"},
	}

	for _, pos := range positions {
		fontSize, exists := fontSizes[pos.Key]
		if !exists {
			fontSize = fontSizes["Default"]
		}
		c.SetFontSize(fontSize)
		if err := drawMultilineText(c, pos.Text, pos.X, pos.Y, lineHeight); err != nil {
			return nil, err
		}
	}

	// Encode the image to a buffer
	buffer := new(bytes.Buffer)
	err = jpeg.Encode(buffer, overlay, nil)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
