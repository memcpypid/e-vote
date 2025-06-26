// module.exports = {
//   content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"], // Tidak ada `purge`
//   theme: {
//     extend: {},
//   },
//   darkMode: "class", // Gunakan 'class' agar bisa dikontrol secara manual
//   plugins: [],
// };
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.html", "./src/**/*.vue"],
  theme: {
    extend: {
      colors: {
        sunny: "#EDEED2",
        black: "#000000",
        sun: "#F56232",
        maryjane: "#AAE73D",
        marydeep: "#79B700",
        ocean: "#186F65",
        fellit: "#E2DFD0",
        deepFilt: "#EEF0E5",
        fillith: "#EEEDEB",
        iceBlue: "#176B87",
        star: "#D2001A",
        che: "#B20600",
        tentrax: "#76885B",
        tentray: "#A3B763",
        tentraC: "#CDE990",
        tentraF: "#BFDB38",
        tentraO: "#9EB23B",
        blues: "#1679AB",
        golkar: "#FFF203",
        golkardeep: "#8c8200",
        nasdem: "#0055FF",
        blues: "#0f0e98",
      },
      fontFamily: {
        Karantina: "Karantina",
        saira: "Saira Condensed",
        Bebas: "Bebas Neue",
        Jet: "JetBrains Mono",
        rubik: "Rubik Mono One",
        anton: "Anton SC",
        cough: "Courier Prime",
      },
      margin: {
        85: "480px",
      },
      width: {
        90: "20rem",
        91: "21rem",
        92: "22.5rem",
        93: "29.5rem",
        95: "28rem",
        98: "34rem",
        100: "32rem",
        105: "45rem",
        110: "50rem",
        120: "60rem",
      },
      height: {
        100: "32rem",
        105: "45rem",
        110: "50rem",
        102: "27rem",
        110: "31rem",
      },
      padding: {
        26: "6.27rem",
        85: "22rem",
        95: "24rem",
        100: "28rem",
        102: "35rem",
        105: "30rem",
        110: "50rem",
      },
    },
    screens: {
      "iphone-11-pro": { raw: "(min-width: 375px) and (max-width: 376px)" },
      siomi: { raw: "(min-width: 374px) and (max-width: 500px)" },
      sm: "640px",

      md: "768px",

      lg: "1024px",

      xl: "1280px",

      "2xl": "1536px",

      "3xl": "1540px",
    },
  },
};
