/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.templ",
    "../../go/pkg/mod/github.com/hbourgeot/templdais@v1.2.1/**/*.templ",
  ],
  theme: {
    extend: {
      height: {
        "screen-page": "calc(100vh - 72px)",
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        lightblue: {
          primary: "#0057D9",
          secondary: "#003C8F",
          accent: "#007BFF",
          neutral: "#1B283F",
          "base-100": "#F1F5F9",
          info: "#1E90FF",
          success: "#4CAF50",
          warning: "#FFC107",
          error: "#DC3545",
        },
      },
      {
        darkblue: {
          primary: "#0057D9",
          secondary: "#021E42",
          accent: "#007BFF",
          neutral: "#101921",
          "base-100": "#121212", // Gris muy oscuro, casi negro, para el fondo
          info: "#1E90FF",
          success: "#198754", // Un verde oscuro para éxitos
          warning: "#FFC107",
          error: "#DC3545",
        },
      },
    ],
  },
};
