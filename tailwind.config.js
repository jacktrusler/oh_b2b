/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['**/*.{html,templ,js}'],
  theme: {
    extend: {
      colors: {
        accent: {
          1: "hsl(var(--color-bkg) / <alpha-value>)",
          2: "hsl(var(--color-bkg) / <alpha-value>)",
        },
        bkg: "hsl(var(--color-bkg) / <alpha-value>)",
        content: "hsl(var(--color-content) / <alpha-value>)",
      },
    },
  },
  plugins: [require('@tailwindcss/forms'), require('@tailwindcss/typography')],
}
