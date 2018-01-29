package shapes

import (
	"fmt"
	"log"
	"math"
	"space"
	"vec3"
)

type Torus struct {
	Position vec3.Vec3
	Radius   float64
	Material space.Material
}

// func (s Torus) Intersect(r ray.Ray) *space.Hit {

//     normal := vec3.Divide(vec3.Subtract(point, s.Position), s.Radius);

//     inclination := math.Acos(normal.Y);
//     azimuth := math.Pi + math.Atan2(normal.X, normal.Z);
//     u := azimuth / (2 * math.Pi);
//     v := inclination / math.Pi;

//     return &space.Hit{
//         T: offset,
//         Position: point,
//         Normal: normal,
//         SurfaceCoordinates: vec3.Vec3{u,v,0},
//         Material: s.Material,
//     }
// }

// double a,b,c,d,e,j,k,l,p,q,t,z;
var a, b, c, d, e, j, k, l, p, q, t, z float64

//explanation of variables
//dy/dx= i* x3 + j*x2 + k*x + l. No i variable needed as it is normalised to 4.
//depressed form of dy/dx: t3+pt+q. A subsitution is used to eliminate the squared term, see below.
//z=q2/4 + p3/27. z is proportional to the "discriminant." The sign tells the number of roots of dy/dx
// +ve z means 1 real root, -ve z means 3 real roots, 0 means one simple root plus a double root.

// recursive newton-raphson to depth i
// void s(double x, int i){
func s(x float64, i int) {
	// double m=(x + b)*x*x*x + c*x*x + d*x + e;
	m := (x+b)*x*x*x + c*x*x + d*x + e
	// if (m == 0 | i == 99)  printf("z= %f x= %f y= %f iteration %d\n", z,x,m,i);
	if m == 0 || i == 99 {
		log.Printf("z= %f x= %f y= %f iteration %d\n", z, x, m, i)
	} else {
		s(x-m/(4*x*x*x+j*x*x+k*k+l), i+1)
	}
	// else s(x-m/(4*x*x*x + j*x*x + k*x + l), i+1);

}

// void _tmain(){
func tmain() {
	// while (true){
	for true {

		// double r[8];
		var r []float64
		//r[1,2,3] and r[5,6,7] store x and y values of the maxima and minima.
		//r[0] and r[4] are dummies to handle [subscript-1] references.

		// scanf_s("%lf %lf %lf %lf %lf", &a, &b, &c, &d, &e);
		//uncomment for verbose output: printf("data for minima and maxima \n");
		//uncomment for verbose output: printf("   z       x       dy/x      y     hi/low \n");
		//hi/low indicate which y values are highest and lowest: 1st, 2nd, 3rd.

		//Divide by a to give a "monic polynomial" with a=1.
		//This does not affect the roots but saves a lot of typing of powers of a.
		// differentiate to i*x3 + j*x2 + k*x + l. As a=1, we know i=4.
		// b /= a; c /= a; d /= a; e /= a;
		b /= a
		c /= a
		d /= a
		e /= a
		// j = 3*b; k = 2*c; l = d;
		j = 3 * b
		k = 2 * c
		l = d

		// to prepare to solve this cubic, convert to depressed cubic t3+pt+q, eliminating x2. x=t-j/(3*i)
		// p=(3*i*k - j2)/(3*i2) q=(2*j3 - 9*i*j*k + 27*i2*l)/27*i3. But we know i=4.
		// p = (12 * k - j*j) / 48;
		p = (12*k - j*j) / 48
		// q = (2 * j*j*j - 36 * j*k + 432 * l) / 1728;
		q = (2*j*j*j - 36*j*k + 432*l) / 1728
		// z = q*q / 4 + p*p*p / 27;
		z = q*q/4 + p*p*p/27

		//u and v indicate which of the stationary points are highest and lowest.
		//solve dy/dx, store the roots in r[1,2,3] and the y values in r[5,6,7].
		// int u=0,v=0,g;
		var u int = 0
		var v int = 0
		var g int
		// for (g = 1; g < 4; g++){
		for g = 1; g < 4; g++ {

			//if z>0 or p==0 use cardano's method to store the one real root of dy/x three times in r[1,2,3]
			//if z<=0 use viete's trigonometric method to find the multiple real roots.
			//use of different methods for each case avoids the need to handle complex numbers.

			// r[g] =z > 0 | p==0? cbrt(-q / 2 + sqrt(z)) + cbrt(-q / 2 - sqrt(z)) - j / 12 :
			//     sqrt(-p / .75)*cos(acos(-q / sqrt(-p * p * p * 4 / 27)) / 3 - g*acos(-.5)) - j / 12;

			if z > 0 || (p == 0) {
				r[g] = math.Cbrt(-q/2+math.Sqrt(z)) + math.Cbrt(-q/2-math.Sqrt(z)) - j/12
			} else {
				r[g] = math.Sqrt(-p/.75)*math.Cos(math.Acos(-q/math.Sqrt(-p*p*p*4/27))/3-float64(g)*math.Acos(-.5)) - j/12
			}

			// r[g+4] = (r[g]+b)*pow(r[g],3) + c*r[g]*r[g] + d*r[g] + e;
			r[g+4] = (r[g]+b)*math.Pow(r[g], 3) + c*r[g]*r[g] + d*r[g] + e

			// if (r[g + 4]>r[g+3]|g==1)u=g;
			if r[g+4] > r[g+3] || float64(g) == 1 {
				u = g
			}
			// if (r[g + 4]<r[g+3]|g==1)v=g;
			if r[g+4] < r[g+3] || float64(g) == 1 {
				v = g
			}

			// uncomment for verbose output
			//printf("%f %f %f %f %d %d \n  \n", z, r[g], 4 * pow(r[g], 3) + j*pow(r[g], 2) + k*r[g] + l, r[g+4], u, v);
		}

		//because we divided by a, the new a=1 and y tends to infinity at large |x|.
		//so if the lowest stationary point has y>0, the whole curve is above the x axis and there is no solution.

		//if the lowest stationary point has y=0 it is tangent to the x axis and would cause problems for newton-raphson

		//if the lowest stationary point has y<0, use newton raphson, 99 iterations. ensure 1st guess is on the side
		//opposite the highest stationary point, just in case that point also has y<0, to avoid getting trapped.
		//special case: if r[v]-r[u] == 0 (implies only one stationary point) then add 1!

		// if (r[v+4] >0) printf("n\n");
		if r[v+4] > 0 {
			log.Print("n\n")
		}
		// if (r[v+4]==0) printf("z= %f x= %f multiple root", z,r[v]);
		if r[v+4] == 0 {
			fmt.Printf("z= %f x= %f multiple root", z, r[v])
		}
		// if (r[v+4] <0) s(r[v]+(r[v]-r[u])+(r[v]-r[u]==0),0);
		if r[v+4] < 0 {
			if r[v]-r[u] == 0 {
				s(r[v]+(r[v]-r[u])+1, 0)
			} else {
				s(r[v]+(r[v]-r[u])+0, 0)
			}
		}
	}
}
