import { ref } from 'vue';
import { eventBus } from '../eventBus.js';

const token = ref(null);

async function login(formData) {
    const response = await fetch('http://localhost:3051/api/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
    });

    if (response.ok) {
        const data = await response.json();
        token.value = data.token;
        eventBus.config.globalProperties.$emit('login', token.value);
    }
}

export { login, token };
