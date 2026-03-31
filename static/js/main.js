document.addEventListener("DOMContentLoaded", function () {
  // Hamburger menu
  var toggle = document.getElementById("nav-toggle");
  var menu = document.getElementById("nav-menu");

  toggle.addEventListener("click", function () {
    menu.classList.toggle("active");
  });

  // Close menu on link click
  var links = document.querySelectorAll(".nav__link");
  links.forEach(function (link) {
    link.addEventListener("click", function () {
      menu.classList.remove("active");
    });
  });

  function createEl(tag, className, text) {
    var el = document.createElement(tag);
    if (className) el.className = className;
    if (text) el.textContent = text;
    return el;
  }

  // Hardcoded data (Phase 1 — will be replaced by API fetch in Phase 2)
  var servicios = [
    { icono: "\uD83D\uDD27", nombre: "Mantenimiento 24hs", descripcion: "Servicio de mantenimiento disponible las 24 horas para cualquier emergencia en tu departamento." },
    { icono: "\uD83D\uDEE1\uFE0F", nombre: "Seguridad", descripcion: "Vigilancia permanente con cámaras de circuito cerrado y personal de seguridad en el edificio." },
    { icono: "\uD83C\uDFCA", nombre: "Áreas Comunes", descripcion: "Piscina, gimnasio, salón de eventos y lavandería comunitaria disponibles para todos los inquilinos." },
    { icono: "\uD83D\uDCB3", nombre: "Administración de Pagos", descripcion: "Gestión centralizada de pagos con múltiples opciones: efectivo, transferencia y billeteras electrónicas." }
  ];

  var departamentos = [
    { nombre: "Monoambiente", precioFormateado: "Gs. 2.500.000/mes", superficie: "35 m²", dormitorios: 0, banos: 1, caracteristicas: ["Cocina integrada", "Balcón", "Aire acondicionado"] },
    { nombre: "1 Dormitorio", precioFormateado: "Gs. 3.500.000/mes", superficie: "55 m²", dormitorios: 1, banos: 1, caracteristicas: ["Cocina separada", "Balcón", "Aire acondicionado", "Placard"] },
    { nombre: "2 Dormitorios", precioFormateado: "Gs. 5.000.000/mes", superficie: "80 m²", dormitorios: 2, banos: 2, caracteristicas: ["Cocina separada", "Balcón amplio", "Aire acondicionado", "2 Placards", "Lavadero"] }
  ];

  // Render servicios
  var serviciosGrid = document.getElementById("servicios-grid");
  servicios.forEach(function (s) {
    var card = createEl("div", "card");
    card.appendChild(createEl("div", "card__icon", s.icono));
    card.appendChild(createEl("h3", "card__title", s.nombre));
    card.appendChild(createEl("p", "card__text", s.descripcion));
    serviciosGrid.appendChild(card);
  });

  // Render departamentos
  var deptGrid = document.getElementById("departamentos-grid");
  departamentos.forEach(function (d) {
    var card = createEl("div", "card");
    card.appendChild(createEl("h3", "card__title", d.nombre));
    card.appendChild(createEl("div", "card__price", d.precioFormateado));

    var meta = d.superficie;
    if (d.dormitorios > 0) {
      meta += " \u00B7 " + d.dormitorios + " dorm.";
    }
    meta += " \u00B7 " + d.banos + " ba\u00F1o" + (d.banos > 1 ? "s" : "");
    card.appendChild(createEl("p", "card__meta", meta));

    var ul = createEl("ul", "card__features");
    d.caracteristicas.forEach(function (c) {
      ul.appendChild(createEl("li", null, c));
    });
    card.appendChild(ul);

    deptGrid.appendChild(card);
  });
});
