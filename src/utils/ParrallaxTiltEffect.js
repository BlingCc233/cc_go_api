export default class ParallaxTiltEffect {
  constructor(element) {
    this.maxRotateDeg = 5;
    this.element = document.querySelectorAll(element)[0];
    this.halfW = 450 / 2;
    this.halfH = 350 / 2;
    this.init();
  }

  init() {
    this.element.addEventListener("mouseenter", this.handleMouseEnter.bind(this));
    this.element.addEventListener("mousemove", this.handleMouseMove.bind(this));
    this.element.addEventListener("mouseleave", this.handleMouseLeave.bind(this));
  }

  computed(offsetX, offsetY) {
    let dxPercent = (offsetX - this.halfW) / this.halfW;
    let dyPercent = -(offsetY - this.halfH) / this.halfH;

    let rotateX = this.maxRotateDeg * dxPercent;
    let rotateY = this.maxRotateDeg * dyPercent;
    this.setElementRotate(rotateY, rotateX);
  }

  handleMouseEnter(e) {
    requestAnimationFrame(() => this.computed(e.offsetX, e.offsetY));
  }

  handleMouseMove(e) {
    requestAnimationFrame(() => this.computed(e.offsetX, e.offsetY));
  }

  handleMouseLeave() {
    this.setElementRotate(0, 0);
  }

  setElementRotate(rotateX, rotateY) {
    this.element.style.setProperty('--X', rotateX + "deg");
    this.element.style.setProperty('--Y', rotateY + "deg");
  }
}
