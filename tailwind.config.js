/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./internal/templates/**/*.templ"],
	theme: {
		extend: {
			colors: {
				"primary-dark": "#403F4C",
				"primary-emerald": "#10b981",
				"primary-orange": "#ea580c",
				bone: "#e3dac9",
				// "pastel-green": "rgb(202, 255, 191)",
				// brown: {
				// 	50: "#fdf8f6",
				// 	100: "#f2e8e5",
				// 	200: "#eaddd7",
				// },
			},
		},
	},
	plugins: [],
};
