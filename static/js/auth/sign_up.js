const ws = new WebSocket("ws://localhost:8080/sign_up_ws");
const form = document.getElementById("authForm");
const statusP = document.getElementById("status");

ws.onopen = () => {
    console.log("WebSocket connected");
};

ws.onmessage = async (event) => {
    const responseText = event.data;
    if (responseText.trim() === "success") {
        statusP.style.color = "green";
        statusP.textContent = "Регистрация успешна! Создание сессии...";
        
        const username = document.getElementById("username").value;

        try {
            const cookieResponse = await fetch("/setcookie", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ username: username })
            });

            if (cookieResponse.ok) {
                statusP.textContent = "Перенаправление...";
                setTimeout(() => {
                    window.location.href = "/";
                }, 1000);
            } else {
                statusP.style.color = "red";
                statusP.textContent = "Ошибка при сохранении сессии";
            }
        } catch (err) {
            console.error("Fetch error:", err);
            statusP.style.color = "red";
            statusP.textContent = "Ошибка сети при установке куки";
        }

    } else {
        statusP.style.color = "red";
        statusP.textContent = "Ошибка: " + responseText;
    }
};

ws.onerror = (error) => {
    console.error("WebSocket error:", error);
    statusP.style.color = "red";
    statusP.textContent = "Ошибка соединения с сервером";
};

const togglePasswordBtn = document.getElementById("togglePassword");
const passwordInput = document.getElementById("password");

if (togglePasswordBtn) {
    togglePasswordBtn.addEventListener("click", () => {
        const type = passwordInput.getAttribute("type") === "password" ? "text" : "password";
        passwordInput.setAttribute("type", type);
        togglePasswordBtn.textContent = type === "password" ? "Показать" : "Скрыть";
    });
}

form.addEventListener("submit", (e) => {
    e.preventDefault();

    const password = passwordInput.value;

    if (password.length < 8) {
        statusP.style.color = "red";
        statusP.textContent = "Ошибка: пароль должен содержать минимум 8 символов";
        return;
    }

    const userData = {
        first_name: document.getElementById("firstName").value,
        last_name: document.getElementById("lastName").value,
        class: document.getElementById("class").value,
        username: document.getElementById("username").value,
        password: password
    };

    if (ws.readyState === WebSocket.OPEN) {
        ws.send(JSON.stringify(userData));
        statusP.style.color = "black";
        statusP.textContent = "Отправка данных...";
    } else {
        statusP.style.color = "red";
        statusP.textContent = "Соединение не установлено";
    }
});