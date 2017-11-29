package random
// 
// const serialVersionUID = 1
// const PHI = 0x9E3779B97F4A7C15
// 
// var x float64
// 
// func InitSeed2(seed float64) {
//     x = seed
// }
// 
// func NewFloat64() {
//     x += PHI
//     return (staffordMix13(x) >> 12 | 0x3FF << 52) - 1.0;
// }
// 
// func staffordMix13(z float64) {
//     z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9;
//     z = (z ^ (z >> 27)) * 0x94D049BB133111EB;
//     return z ^ (z >> 31);
// }