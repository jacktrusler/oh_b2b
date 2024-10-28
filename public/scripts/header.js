document.addEventListener("DOMContentLoaded", () => {
  const loginButton = document.getElementById("login-button");
  const signupButton = document.getElementById("signup-button");
  const loginDialog = document.getElementById("login-dialog");
  const signupDialog = document.getElementById("signup-dialog");
  const loginCloseButton = document.getElementById("login-close-button");
  const signupCloseButton = document.getElementById("signup-close-button");

  loginButton.addEventListener("click", () => {
    loginDialog.showModal();
  });

  signupButton.addEventListener("click", () => {
    signupDialog.showModal();
  });

  loginCloseButton.addEventListener("click", () => {
    loginDialog.close();
  });

  signupCloseButton.addEventListener("click", () => {
    signupDialog.close();
  });
})
