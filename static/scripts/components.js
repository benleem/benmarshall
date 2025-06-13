class TextArea extends HTMLElement {
	form = document.getElementById("contact-form");
	input = document.getElementById("contact-message");

	constructor() {
		super();

		this.resize(this.input);
	}

	connectedCallback() {
		this.input.addEventListener("input", (event) => {
			this.resize(event.target);
		});

		this.form.addEventListener("reset", () => {
			this.formReset();
		});
		// htmx.process(this);
	}

	formReset = () => {
		this.input.style.height = "min-content";
	};

	resize = (target) => {
		if (target.value.length === 0) {
			target.style.height = "auto";
		} else {
			target.style.height = "auto";
			target.style.height = `${target.scrollHeight}px`;
			target.style.overflowY = "hidden";
		}
	};

	disconnectedCallback() {
		this.input.removeEventListener("input", this.resize);
		this.form.removeEventListener("reset", this.formReset);
	}
}

customElements.define("text-area", TextArea);
