document.addEventListener("DOMContentLoaded", function () {
  var toggleBtn = document.getElementById("chatbot-toggle");
  var closeBtn = document.getElementById("chatbot-close");
  var chatWindow = document.getElementById("chatbot-window");
  var messagesEl = document.getElementById("chatbot-messages");
  var form = document.getElementById("chatbot-form");
  var input = document.getElementById("chatbot-input");

  var faqs = [
    { keywords: ["horario", "hora", "atencion", "abierto", "abre", "cierra"], respuesta: "Nuestro horario de atención es de lunes a viernes de 08:00 a 17:00 y sábados de 08:00 a 12:00." },
    { keywords: ["ubicacion", "direccion", "donde", "llegar", "mapa"], respuesta: "Estamos ubicados en Av. Mariscal López 1234, Asunción, Paraguay." },
    { keywords: ["precio", "alquiler", "costo", "cuanto", "mensualidad", "renta"], respuesta: "Nuestros departamentos van desde Gs. 2.500.000/mes (monoambiente) hasta Gs. 5.000.000/mes (2 dormitorios). Contactanos para conocer disponibilidad actual." },
    { keywords: ["pago", "pagar", "transferencia", "efectivo", "tarjeta", "deposito"], respuesta: "Aceptamos pago en efectivo, transferencia bancaria y giros desde billeteras electrónicas. El pago se realiza del 1 al 5 de cada mes." },
    { keywords: ["mantenimiento", "reparacion", "arreglo", "roto", "problema", "fuga", "luz"], respuesta: "Para solicitudes de mantenimiento, podés comunicarte al (021) 555-1234 o enviar un mensaje por este chat. Atendemos urgencias las 24 horas." },
    { keywords: ["requisito", "documento", "contrato", "garantia", "alquilar", "inquilino"], respuesta: "Para alquilar necesitás: cédula de identidad, comprobante de ingresos, referencia personal y un depósito de garantía equivalente a 2 meses de alquiler." },
    { keywords: ["area", "comun", "piscina", "gimnasio", "estacionamiento", "salon", "lavanderia"], respuesta: "Contamos con piscina, gimnasio, salón de eventos, lavandería comunitaria y estacionamiento para inquilinos." },
    { keywords: ["contacto", "telefono", "llamar", "email", "correo", "whatsapp"], respuesta: "Podés contactarnos al (021) 555-1234, por WhatsApp al +595 981 123456, o por correo a info@residencialguarani.com.py." },
    { keywords: ["mascota", "perro", "gato", "animal"], respuesta: "Se permiten mascotas pequeñas (hasta 10 kg) con un depósito adicional de Gs. 500.000. Consultá las normas de convivencia en administración." },
    { keywords: ["incluido", "servicio", "agua", "internet", "basura", "expensa", "gasto"], respuesta: "El alquiler incluye agua, recolección de basura y mantenimiento de áreas comunes. Internet y electricidad corren por cuenta del inquilino." }
  ];

  var easterEggs = [
    { keywords: ["urgencia", "urgencias"], respuesta: "Jaja si atendemos 😄" }
  ];

  var greetings = ["hola", "buenas", "buenos", "buen dia", "buenas tardes", "buenas noches"];
  var fallback = "No entendí tu consulta, ¿podés reformularla? Puedo ayudarte con horarios, precios, ubicación, mantenimiento y más.";

  function normalize(text) {
    return text
      .toLowerCase()
      .normalize("NFD")
      .replace(/[\u0300-\u036f]/g, "")
      .trim();
  }

  function addMessage(text, sender) {
    var msg = document.createElement("div");
    msg.className = "chatbot__msg chatbot__msg--" + sender;
    msg.textContent = text;
    messagesEl.appendChild(msg);
    messagesEl.scrollTop = messagesEl.scrollHeight;
  }

  function getResponse(userText) {
    var normalized = normalize(userText);

    for (var i = 0; i < greetings.length; i++) {
      if (normalized.indexOf(greetings[i]) !== -1) {
        return "\u00A1Hola! Soy el asistente virtual de Residencial Guaraní. \u00BFEn qué puedo ayudarte?";
      }
    }

    for (var e = 0; e < easterEggs.length; e++) {
      for (var ek = 0; ek < easterEggs[e].keywords.length; ek++) {
        if (normalized.indexOf(easterEggs[e].keywords[ek]) !== -1) {
          return easterEggs[e].respuesta;
        }
      }
    }

    var words = normalized.split(/\s+/);
    var bestMatch = null;
    var bestScore = 0;

    for (var f = 0; f < faqs.length; f++) {
      var score = 0;
      for (var k = 0; k < faqs[f].keywords.length; k++) {
        for (var w = 0; w < words.length; w++) {
          if (words[w].indexOf(faqs[f].keywords[k]) !== -1) {
            score++;
            break;
          }
        }
      }
      if (score > bestScore) {
        bestScore = score;
        bestMatch = faqs[f];
      }
    }

    if (bestScore > 0 && bestMatch) {
      return bestMatch.respuesta;
    }
    return fallback;
  }

  // Toggle chat window
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

  // Handle form submit
  form.addEventListener("submit", function (e) {
    e.preventDefault();
    var text = input.value.trim();
    if (!text) return;

    addMessage(text, "user");
    input.value = "";

    var response = getResponse(text);
    setTimeout(function () {
      addMessage(response, "bot");
    }, 400);
  });

  // Welcome message
  addMessage("\u00A1Hola! Soy el asistente virtual de Residencial Guaraní. Podés preguntarme sobre horarios, precios, ubicación, mantenimiento y más.", "bot");
});
