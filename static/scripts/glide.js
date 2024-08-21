var glide = new Glide(".glide", {
	type: "carousel",
	autoplay: 3000,
	keyboard: true,
	startAt: 0,
	focusAt: "center",
	perView: 3,
	peek: {
		before: 0,
		after: 0,
	},
	gap: 20,
	breakpoints: {
		1200: {
			perView: 2,
		},
		992: {
			perView: 1,
		},
	},
});
glide.mount();

// function resize() {
// 	const { innerWidth } = window;
// 	nav = document.getElementById("nav-aside");
// 	main = document.querySelector("main");
// 	main.style.width = `${innerWidth - nav.offsetWidth - 100}px`;
// }
// resize();
// window.addEventListener("resize", resize);
