const login_container = document.querySelector('.login-container');

function setCookie(name, value) {
    document.cookie = `${name}=${value}; path=/`;
}

if (login_container) {
    const emailInput = document.querySelector('#login-email');
    const passwordInput = document.querySelector('#login-password');
    const loginButton = document.querySelector('#login-submit');

    loginButton.addEventListener('click', async () => {
        const email = emailInput.value;
        const password = passwordInput.value;

        if (!email || !password) {
            alert('Please enter both email and password.');
            return;
        }

        try {
            const response = await fetch('/api/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ 
                    'email': email, 
                    'password': password 
                }),
            });

            if (response.ok) {
                const data = await response.json();
                const token = data.token;
                const id = data.id;
                alert('Login response:', data);
                if (token && id) {
                    
                    setCookie('session_token', token);
                    setCookie('session_user_id', id);
                    alert('Login successful!');
                }
            } else {
                const errorData = await response.json();
                alert('Login failed: ' + errorData.error);
            }
        } catch (error) {
            console.error('Error during login:', error);
            alert('An error occurred. Please try again later.');
        }
    });
}