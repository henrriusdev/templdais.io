/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.{templ,go}",
    "../../go/pkg/mod/github.com/hbourgeot/templdais@v1.0.2-0.20240428174538-7f41dd2a6410/**/*.{templ,go}",
  ],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    theme: ["night", "light"],
  },
};

