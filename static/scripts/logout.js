const logout_button = document.querySelector('#logout-button');

logout_button.addEventListener('click', async () => {
    document.cookie = 'session_token=;path=/';
    document.cookie = 'session_user_id=;path=/';
});