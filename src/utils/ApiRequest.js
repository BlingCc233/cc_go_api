async function makeApiRequest(apiData) {
    try {
        const url = new URL(apiData.value.apiurl);

        // 如果是 GET 请求，将参数附加到 URL
        if (apiData.value.method.toUpperCase() === 'GET') {
            for (const [key, value] of Object.entries(apiData.value.params)) {
                url.searchParams.append(key, value);
            }
        }

        // 配置请求选项
        const options = {
            method: apiData.value.method,
            headers: {
                'Content-Type': 'application/json', // 根据需要设置请求头
                ...apiData.value.headers // 如果有其他自定义头部
            },
            body: apiData.value.method.toUpperCase() !== 'GET' ? JSON.stringify(apiData.value.body) : undefined // 仅在非 GET 请求时设置 body
        };

        const res = await fetch(url, options);

        // 检查响应状态
        if (!res.ok) {
            throw new Error(`HTTP error! status: ${res.status}`);
        }

        // 返回响应数据
        return await res.json();
    } catch (error) {
        console.error('Error making API request:', error);
        throw error; // 重新抛出错误以便调用者处理
    }
}

export default makeApiRequest;
