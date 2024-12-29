import daisyui from "daisyui"
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["public/views/**/*.{templ,js}"],
  theme: {
    extend: {},
  },
  plugins: [
    daisyui,
  ],
}
