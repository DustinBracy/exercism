export class Triangle {
  constructor(...sides) {
    this.sides = [...sides]
    this.a = sides[0]
    this.b = sides[1]
    this.c = sides[2]
  }

  get isValid() {
    return this.sides.every(side => side > 0) && (this.a + this.b >= this.c && this.a + this.c >= this.b && this.c + this.b >= this.a)
  }

  get isEquilateral() {
    return this.isValid && this.a === this.b && this.a === this.c
  }

  get isIsosceles() {
    return this.isValid && (this.a === this.b || this.a === this.c || this.b === this.c)
  }

  get isScalene() {
    return this.isValid && this.a != this.b && this.a != this.c && this.b != this.c
  }
}
