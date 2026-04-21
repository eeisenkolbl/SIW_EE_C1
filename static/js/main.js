document.addEventListener("DOMContentLoaded", function () {
  // Hamburger menu
  var toggle = document.getElementById("nav-toggle");
  var menu = document.getElementById("nav-menu");

  toggle.addEventListener("click", function () {
    menu.classList.toggle("active");
  });

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

  function renderServicios(servicios) {
    var grid = document.getElementById("servicios-grid");
    servicios.forEach(function (s) {
      var card = createEl("div", "card");
      card.appendChild(createEl("div", "card__icon", s.icono));
      card.appendChild(createEl("h3", "card__title", s.nombre));
      card.appendChild(createEl("p", "card__text", s.descripcion));
      grid.appendChild(card);
    });
  }

  function renderDepartamentos(departamentos) {
    var grid = document.getElementById("departamentos-grid");
    departamentos.forEach(function (d) {
      var card = createEl("div", "card");

      if (d.imagen) {
        var img = document.createElement("img");
        img.src = d.imagen;
        img.alt = d.nombre;
        img.className = "card__image";
        img.loading = "lazy";
        card.appendChild(img);
      }

      card.appendChild(createEl("h3", "card__title", d.nombre));
      card.appendChild(createEl("div", "card__price", d.precioFormateado));

      var meta = d.superficie;
      if (d.dormitorios > 0) {
        meta += " · " + d.dormitorios + " dorm.";
      }
      meta += " · " + d.banos + " baño" + (d.banos > 1 ? "s" : "");
      card.appendChild(createEl("p", "card__meta", meta));

      var ul = createEl("ul", "card__features");
      d.caracteristicas.forEach(function (c) {
        ul.appendChild(createEl("li", null, c));
      });
      card.appendChild(ul);

      grid.appendChild(card);
    });
  }

  fetch("/api/servicios")
    .then(function (r) { return r.json(); })
    .then(renderServicios);

  fetch("/api/departamentos")
    .then(function (r) { return r.json(); })
    .then(renderDepartamentos);
});
