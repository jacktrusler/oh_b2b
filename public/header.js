const signInButton = document.getElementById("sign-in-button");
const loginForm = document.getElementById("login-form");
const signInDialog = document.getElementById("sign-in-dialog");
const closeButton = document.getElementById("close-button");

signInButton.addEventListener("click", () => {
  signInDialog.showModal();
});

closeButton.addEventListener("click", () => {
  signInDialog.close();
});

loginForm.addEventListener("submit", async (e) => {
  const text = document.getElementById("submit-text");
  const loginButton = document.getElementById("login-button");
  const buttonSpinner = document.getElementById("button-spinner");
  const message = document.getElementById("login-message");
  e.preventDefault();

  text.classList.add("hidden")
  buttonSpinner.classList.add("spinner-show");
  loginButton.disabled = true;

  const formData = new FormData(loginForm);

  try {
    const res = await fetch("/login", {
      method: "POST",
      body: formData,
    });

    buttonSpinner.classList.remove("spinner-show");
    text.classList.remove("hidden");
    loginButton.disabled = false;

    console.log(res)

    if (res.ok) {
      message.textContent = "Login was successful!"
    } else {
      message.textContent = "Login was unsuccessful"
    }

  } catch (err) {

  }
});
