/* ----------------------------
 Toggling Dark and Light mode
 ---------------------------- */
const toggle = document.getElementById("theme-toggle");

const storedTheme = window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
if (storedTheme) {
  document.documentElement.setAttribute('data-theme', storedTheme);
}

const dT = document.documentElement.getAttribute("data-theme");
const sun = document.getElementById("sun");
const moon = document.getElementById("moon");
const logoPic = document.getElementById("logo-pic");


if (dT === "dark") {
  moon.setAttribute("class", "moon hidden");
  logoPic.setAttribute("src", "/public/assets/b2blight.png");
}
if (dT === "light") {
  sun.setAttribute("class", "sun hidden");
  logoPic.setAttribute("src", "/public/assets/b2blogo.png");
}


if (toggle) {
  toggle.onclick = function() {
    const currentTheme = document.documentElement.getAttribute("data-theme");
    let targetTheme = "light";

    if (currentTheme === "light") {
      targetTheme = "dark";
      moon.setAttribute("class", "moon hidden")
      sun.setAttribute("class", "sun");
      logoPic.setAttribute("src", "/public/assets/b2blight.png");
    } else {
      sun.setAttribute("class", "sun hidden");
      moon.setAttribute("class", "moon");
      logoPic.setAttribute("src", "/public/assets/b2blogo.png");
    }

    document.documentElement.setAttribute('data-theme', targetTheme);
    localStorage.setItem('theme', targetTheme);
  }
}
