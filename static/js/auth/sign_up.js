const ws = new WebSocket("ws://localhost:8080/sign_in_ws");
const form = document.getElementById("authForm");
const statusP = document.getElementById("status");

ws.onopen = () => {
    console.log("WebSocket connected");
};

ws.onmessage = (event) => {
    const responseText = event.data;
    if (responseText.trim() === "success") {
        statusP.style.color = "green";
        statusP.textContent = "Регистрация успешна! Перенаправление...";
        
        setTimeout(() => {
            window.location.href = "/";
        }, 1500);
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

form.addEventListener("submit", (e) => {
    e.preventDefault();

    const userData = {
        first_name: document.getElementById("firstName").value,
        last_name: document.getElementById("lastName").value,
        class: document.getElementById("class").value,
        username: document.getElementById("username").value,
        password: document.getElementById("password").value
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