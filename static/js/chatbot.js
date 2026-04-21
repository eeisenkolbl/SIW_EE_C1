document.addEventListener("DOMContentLoaded", function () {
  var toggleBtn = document.getElementById("chatbot-toggle");
  var closeBtn = document.getElementById("chatbot-close");
  var chatWindow = document.getElementById("chatbot-window");
  var messagesEl = document.getElementById("chatbot-messages");
  var form = document.getElementById("chatbot-form");
  var input = document.getElementById("chatbot-input");

  function addMessage(text, sender) {
    var msg = document.createElement("div");
    msg.className = "chatbot__msg chatbot__msg--" + sender;
    msg.textContent = text;
    messagesEl.appendChild(msg);
    messagesEl.scrollTop = messagesEl.scrollHeight;
  }

  toggleBtn.addEventListener("click", function () {
    var isHidden = chatWindow.hasAttribute("hidden");
    if (isHidden) {
      chatWindow.removeAttribute("hidden");
      input.focus();
    } else {
      chatWindow.setAttribute("hidden", "");
    }
  });

  closeBtn.addEventListener("click", function () {
    chatWindow.setAttribute("hidden", "");
  });

  form.addEventListener("submit", function (e) {
    e.preventDefault();
    var text = input.value.trim();
    if (!text) return;

    addMessage(text, "user");
    input.value = "";

    fetch("/api/chat", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ message: text })
    })
      .then(function (r) { return r.json(); })
      .then(function (data) { addMessage(data.response, "bot"); })
      .catch(function () { addMessage("Error al conectar con el servidor.", "bot"); });
  });

  addMessage("¡Hola! Soy el asistente virtual de Residencial Guaraní. Podés preguntarme sobre horarios, precios, ubicación, mantenimiento y más.", "bot");
});
