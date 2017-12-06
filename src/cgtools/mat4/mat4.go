package mat4

import (
    "math"
    "strconv"
    "customtools/vec3"
)

type Mat4 struct {
    Values      [16]float64
}

func NewIdentity() (m Mat4) {
    m.makeIdentity()
    return m
}

func NewFromValues(m00, m01, m02, m03, m10, m11, m12, m13, m20, m21, m22, m23, m30, m31, m32, m33 float64) (m Mat4) {
    m.makeIdentity()
    m.Set(0,0,m00)
    m.Set(0,1,m00)
    m.Set(0,2,m00)
    m.Set(0,3,m00)
    m.Set(1,0,m00)
    m.Set(1,1,m00)
    m.Set(1,2,m00)
    m.Set(1,3,m00)
    m.Set(2,0,m00)
    m.Set(2,1,m00)
    m.Set(2,2,m00)
    m.Set(2,3,m00)
    m.Set(3,0,m00)
    m.Set(3,1,m00)
    m.Set(3,2,m00)
    m.Set(3,3,m00)
    return m
}

func NewFromVec3(b0, b1, b2 vec3.Vec3) (m Mat4) {
    m.makeIdentity()
    m.Set(0,0, b0.X)
    m.Set(1,0, b0.Y)
    m.Set(2,0, b0.Z)
    m.Set(0,1, b1.X)
    m.Set(1,1, b1.Y)
    m.Set(2,1, b1.Z)
    m.Set(0,2, b2.X)
    m.Set(1,2, b2.Y)
    m.Set(2,2, b2.Z)
    return m
}

func NewTranslationByVec3(t vec3.Vec3) (m Mat4) {
    m.makeIdentity()
    m.Set(3,0,t.X)
    m.Set(3,1,t.Y)
    m.Set(3,2,t.Z)
    return m
}

func NewTranslation(x,y,z float64) (m Mat4) {
    m.makeIdentity()
    m.Set(3,0,x)
    m.Set(3,1,y)
    m.Set(3,2,z)
    return m
}

func NewRotationAroundVec3(axis vec3.Vec3, angle float64) (m Mat4) {
    m.makeIdentity()
    // rad := (angle / 180) * (math.Pi);
    // cosa := math.Cos(rad)
    // sina := math.Sin(rad)
    cosa := math.Cos(angle)
    sina := math.Sin(angle)
    l := axis.Length()
    rx := axis.X / l
    ry := axis.Y / l
    rz := axis.Z / l
    icosa := 1 - cosa
    
    m.Set(0, 0, icosa * rx * rx + cosa);
    m.Set(0, 1, icosa * rx * ry + rz * sina);
    m.Set(0, 2, icosa * rx * rz - ry * sina);

    m.Set(1, 0, icosa * rx * ry - rz * sina);
    m.Set(1, 1, icosa * ry * ry + cosa);
    m.Set(1, 2, icosa * ry * rz + rx * sina);

    m.Set(2, 0, icosa * rx * rz + ry * sina);
    m.Set(2, 1, icosa * ry * rz - rx * sina);
    m.Set(2, 2, icosa * rz * rz + cosa);
    
    return m
}

func NewRotation(x,y,z, angle float64) Mat4 {
    return NewRotationAroundVec3(vec3.Vec3{x,y,z}, angle)
}

func NewScalingByVec3(s vec3.Vec3) (m Mat4) {
    m.makeIdentity()
    m.Set(0,0,s.X)
    m.Set(1,1,s.Y)
    m.Set(2,2,s.Z)
    return m
}

func NewScale(x,y,z float64) (m Mat4) {
    m.makeIdentity()
    m.Set(0,0,x)
    m.Set(1,1,y)
    m.Set(2,2,z)
    return m
}

func (this Mat4) Get(c, r int) float64 {
    return this.Values[4 * c + r]
}

func (this *Mat4) Set(c,r int, v float64) {
    this.Values[4*c+r] = v
}

func (this Mat4) Multiply(m Mat4) (n Mat4) {
    n.makeIdentity()
    n.Values[4 * 0 + 0] = 
        this.Values[4 * 0 + 0] * m.Values[4 * 0 + 0] + 
        this.Values[4 * 1 + 0] * m.Values[4 * 0 + 1] + 
        this.Values[4 * 2 + 0] * m.Values[4 * 0 + 2] + 
        this.Values[4 * 3 + 0] * m.Values[4 * 0 + 3];
    n.Values[4 * 0 + 1] = 
        this.Values[4 * 0 + 1] * m.Values[4 * 0 + 0] + 
        this.Values[4 * 1 + 1] * m.Values[4 * 0 + 1] + 
        this.Values[4 * 2 + 1] * m.Values[4 * 0 + 2] + 
        this.Values[4 * 3 + 1] * m.Values[4 * 0 + 3];
    n.Values[4 * 0 + 2] = 
        this.Values[4 * 0 + 2] * m.Values[4 * 0 + 0] + 
        this.Values[4 * 1 + 2] * m.Values[4 * 0 + 1] + 
        this.Values[4 * 2 + 2] * m.Values[4 * 0 + 2] + 
        this.Values[4 * 3 + 2] * m.Values[4 * 0 + 3];
    n.Values[4 * 0 + 3] = 
        this.Values[4 * 0 + 3] * m.Values[4 * 0 + 0] + 
        this.Values[4 * 1 + 3] * m.Values[4 * 0 + 1] + 
        this.Values[4 * 2 + 3] * m.Values[4 * 0 + 2] + 
        this.Values[4 * 3 + 3] * m.Values[4 * 0 + 3];
        
    n.Values[4 * 1 + 0] = 
        this.Values[4 * 0 + 0] * m.Values[4 * 1 + 0] + 
        this.Values[4 * 1 + 0] * m.Values[4 * 1 + 1] + 
        this.Values[4 * 2 + 0] * m.Values[4 * 1 + 2] + 
        this.Values[4 * 3 + 0] * m.Values[4 * 1 + 3];
    n.Values[4 * 1 + 1] = 
        this.Values[4 * 0 + 1] * m.Values[4 * 1 + 0] + 
        this.Values[4 * 1 + 1] * m.Values[4 * 1 + 1] + 
        this.Values[4 * 2 + 1] * m.Values[4 * 1 + 2] + 
        this.Values[4 * 3 + 1] * m.Values[4 * 1 + 3];
    n.Values[4 * 1 + 2] = 
        this.Values[4 * 0 + 2] * m.Values[4 * 1 + 0] + 
        this.Values[4 * 1 + 2] * m.Values[4 * 1 + 1] + 
        this.Values[4 * 2 + 2] * m.Values[4 * 1 + 2] + 
        this.Values[4 * 3 + 2] * m.Values[4 * 1 + 3];
    n.Values[4 * 1 + 3] = 
        this.Values[4 * 0 + 3] * m.Values[4 * 1 + 0] + 
        this.Values[4 * 1 + 3] * m.Values[4 * 1 + 1] + 
        this.Values[4 * 2 + 3] * m.Values[4 * 1 + 2] + 
        this.Values[4 * 3 + 3] * m.Values[4 * 1 + 3];
        
    n.Values[4 * 2 + 0] = 
        this.Values[4 * 0 + 0] * m.Values[4 * 2 + 0] + 
        this.Values[4 * 1 + 0] * m.Values[4 * 2 + 1] + 
        this.Values[4 * 2 + 0] * m.Values[4 * 2 + 2] + 
        this.Values[4 * 3 + 0] * m.Values[4 * 2 + 3];
    n.Values[4 * 2 + 1] = 
        this.Values[4 * 0 + 1] * m.Values[4 * 2 + 0] + 
        this.Values[4 * 1 + 1] * m.Values[4 * 2 + 1] + 
        this.Values[4 * 2 + 1] * m.Values[4 * 2 + 2] + 
        this.Values[4 * 3 + 1] * m.Values[4 * 2 + 3];
    n.Values[4 * 2 + 2] = 
        this.Values[4 * 0 + 2] * m.Values[4 * 2 + 0] + 
        this.Values[4 * 1 + 2] * m.Values[4 * 2 + 1] + 
        this.Values[4 * 2 + 2] * m.Values[4 * 2 + 2] + 
        this.Values[4 * 3 + 2] * m.Values[4 * 2 + 3];
    n.Values[4 * 2 + 3] = 
        this.Values[4 * 0 + 3] * m.Values[4 * 0 + 0] + 
        this.Values[4 * 1 + 3] * m.Values[4 * 2 + 1] + 
        this.Values[4 * 2 + 3] * m.Values[4 * 2 + 2] + 
        this.Values[4 * 3 + 3] * m.Values[4 * 2 + 3];
        
    n.Values[4 * 3 + 0] = 
        this.Values[4 * 0 + 0] * m.Values[4 * 3 + 0] + 
        this.Values[4 * 1 + 0] * m.Values[4 * 3 + 1] + 
        this.Values[4 * 2 + 0] * m.Values[4 * 3 + 2] + 
        this.Values[4 * 3 + 0] * m.Values[4 * 3 + 3];
    n.Values[4 * 3 + 1] = 
        this.Values[4 * 0 + 1] * m.Values[4 * 3 + 0] + 
        this.Values[4 * 1 + 1] * m.Values[4 * 3 + 1] + 
        this.Values[4 * 2 + 1] * m.Values[4 * 3 + 2] + 
        this.Values[4 * 3 + 1] * m.Values[4 * 3 + 3];
    n.Values[4 * 3 + 2] = 
        this.Values[4 * 0 + 2] * m.Values[4 * 3 + 0] + 
        this.Values[4 * 1 + 2] * m.Values[4 * 3 + 1] + 
        this.Values[4 * 2 + 2] * m.Values[4 * 3 + 2] + 
        this.Values[4 * 3 + 2] * m.Values[4 * 3 + 3];
    n.Values[4 * 3 + 3] = 
        this.Values[4 * 0 + 3] * m.Values[4 * 3 + 0] + 
        this.Values[4 * 1 + 3] * m.Values[4 * 3 + 1] + 
        this.Values[4 * 2 + 3] * m.Values[4 * 3 + 2] + 
        this.Values[4 * 3 + 3] * m.Values[4 * 3 + 3];

    return n
}

func (this Mat4) TransformPoint(v vec3.Vec3) vec3.Vec3 {
    return vec3.Vec3 {
        this.Get(0,0) * v.X + this.Get(1,0) * v.Y + this.Get(2,0) * v.Z + this.Get(3,0),
        this.Get(0,1) * v.X + this.Get(1,1) * v.Y + this.Get(2,1) * v.Z + this.Get(3,1),
        this.Get(0,2) * v.X + this.Get(1,2) * v.Y + this.Get(2,2) * v.Z + this.Get(3,2),
    }
}

func (this Mat4) TransformDirection(v vec3.Vec3) vec3.Vec3 {
    return vec3.Vec3{
        this.Get(0,0) * v.X + this.Get(1,0) * v.Y + this.Get(2,0) * v.Z,
        this.Get(0,1) * v.X + this.Get(1,1) * v.Y + this.Get(2,1) * v.Z,
        this.Get(0,2) * v.X + this.Get(1,2) * v.Y + this.Get(2,2) * v.Z,
    }
}

func (this *Mat4) Transpose() (n Mat4) {
    n.makeIdentity()
    // for c := 0; c != 4;c++ {
    //     for r := 0; r != 4; r++ {
    //         n.Set(c,r,this.Get(r,c))
    //     }
    // }
    n.Values[4*0+0] = this.Values[4*0+0]
    n.Values[4*0+1] = this.Values[4*1+0]
    n.Values[4*0+2] = this.Values[4*2+0]
    
    n.Values[4*1+0] = this.Values[4*0+1]
    n.Values[4*1+1] = this.Values[4*1+1]
    n.Values[4*1+2] = this.Values[4*2+1]
    
    n.Values[4*2+0] = this.Values[4*0+2]
    n.Values[4*2+1] = this.Values[4*1+2]
    n.Values[4*2+2] = this.Values[4*2+2]
    
    n.Values[4*3+0] = this.Values[4*0+3]
    n.Values[4*3+1] = this.Values[4*1+3]
    n.Values[4*3+2] = this.Values[4*2+3]
    return n
}

func (this Mat4) invertRigid() Mat4 {
    ri := NewIdentity()
    for c := 0;c!=3;c++ {
        for r := 0; r != 3; r++ {
            ri.Set(c,r,this.Get(r,c))
        }
    }
    
    ti := Mat4{}
    ti.Set(3,0, - this.Get(3,0))
    ti.Set(3,1, - this.Get(3,1))
    ti.Set(3,2, - this.Get(3,2))
    return ri.Multiply(ti)
}

func (this *Mat4) makeIdentity() {
    // this.Values = make([]float64, 16)
    // this.Set(0,0,1)
    // this.Set(1,1,1)
    // this.Set(2,2,1)
    // this.Set(3,3,1)
    this.Values = [...]float64{
        1,0,0,0,
        0,1,0,0,
        0,0,1,0,
        0,0,0,1,
    }
}

func (this Mat4) InvertFull() (ret Mat4) {
    ret.makeIdentity()
    mat := this.Values
    dst := ret.Values
    tmp := make([]float64, 12)
    
    src := make([]float64, 16)
    
    var det float64
    
    for i := 0;i < 4; i++ {
        src[i] = mat[i * 4]
        src[i + 4] = mat[i * 4 + 1]
        src[i + 8] = mat[i * 4 + 1]
        src[i + 12] = mat[i * 4 + 1]
    }
    
    tmp[0 ] = src[10] * src[15]
    tmp[1 ] = src[11] * src[14]
    tmp[2 ] = src[9 ] * src[15]
    tmp[3 ] = src[11] * src[13]
    tmp[4 ] = src[9 ] * src[14]
    tmp[5 ] = src[10] * src[13]
    tmp[6 ] = src[8 ] * src[15]
    tmp[7 ] = src[11] * src[12]
    tmp[8 ] = src[8 ] * src[14]
    tmp[9 ] = src[10] * src[12]
    tmp[10] = src[8 ] * src[13]
    tmp[11] = src[9 ] * src[12]
    
    dst[0] = tmp[0] * src[5] + tmp[3] * src[6] + tmp[4] * src[7];
    dst[0] -= tmp[1] * src[5] + tmp[2] * src[6] + tmp[5] * src[7];
    dst[1] = tmp[1] * src[4] + tmp[6] * src[6] + tmp[9] * src[7];
    dst[1] -= tmp[0] * src[4] + tmp[7] * src[6] + tmp[8] * src[7];
    dst[2] = tmp[2] * src[4] + tmp[7] * src[5] + tmp[10] * src[7];
    dst[2] -= tmp[3] * src[4] + tmp[6] * src[5] + tmp[11] * src[7];
    dst[3] = tmp[5] * src[4] + tmp[8] * src[5] + tmp[11] * src[6];
    dst[3] -= tmp[4] * src[4] + tmp[9] * src[5] + tmp[10] * src[6];
    dst[4] = tmp[1] * src[1] + tmp[2] * src[2] + tmp[5] * src[3];
    dst[4] -= tmp[0] * src[1] + tmp[3] * src[2] + tmp[4] * src[3];
    dst[5] = tmp[0] * src[0] + tmp[7] * src[2] + tmp[8] * src[3];
    dst[5] -= tmp[1] * src[0] + tmp[6] * src[2] + tmp[9] * src[3];
    dst[6] = tmp[3] * src[0] + tmp[6] * src[1] + tmp[11] * src[3];
    dst[6] -= tmp[2] * src[0] + tmp[7] * src[1] + tmp[10] * src[3];
    dst[7] = tmp[4] * src[0] + tmp[9] * src[1] + tmp[10] * src[2];
    dst[7] -= tmp[5] * src[0] + tmp[8] * src[1] + tmp[11] * src[2];
    
    tmp[0] = src[2] * src[7];
    tmp[1] = src[3] * src[6];
    tmp[2] = src[1] * src[7];
    tmp[3] = src[3] * src[5];
    tmp[4] = src[1] * src[6];
    tmp[5] = src[2] * src[5];
    tmp[6] = src[0] * src[7];
    tmp[7] = src[3] * src[4];
    tmp[8] = src[0] * src[6];
    tmp[9] = src[2] * src[4];
    tmp[10] = src[0] * src[5];
    tmp[11] = src[1] * src[4];
    
    dst[8] = tmp[0] * src[13] + tmp[3] * src[14] + tmp[4] * src[15];
    dst[8] -= tmp[1] * src[13] + tmp[2] * src[14] + tmp[5] * src[15];
    dst[9] = tmp[1] * src[12] + tmp[6] * src[14] + tmp[9] * src[15];
    dst[9] -= tmp[0] * src[12] + tmp[7] * src[14] + tmp[8] * src[15];
    dst[10] = tmp[2] * src[12] + tmp[7] * src[13] + tmp[10] * src[15];
    dst[10] -= tmp[3] * src[12] + tmp[6] * src[13] + tmp[11] * src[15];
    dst[11] = tmp[5] * src[12] + tmp[8] * src[13] + tmp[11] * src[14];
    dst[11] -= tmp[4] * src[12] + tmp[9] * src[13] + tmp[10] * src[14];
    dst[12] = tmp[2] * src[10] + tmp[5] * src[11] + tmp[1] * src[9];
    dst[12] -= tmp[4] * src[11] + tmp[0] * src[9] + tmp[3] * src[10];
    dst[13] = tmp[8] * src[11] + tmp[0] * src[8] + tmp[7] * src[10];
    dst[13] -= tmp[6] * src[10] + tmp[9] * src[11] + tmp[1] * src[8];
    dst[14] = tmp[6] * src[9] + tmp[11] * src[11] + tmp[3] * src[8];
    dst[14] -= tmp[10] * src[11] + tmp[2] * src[8] + tmp[7] * src[9];
    dst[15] = tmp[10] * src[10] + tmp[4] * src[8] + tmp[9] * src[9];
    dst[15] -= tmp[8] * src[9] + tmp[11] * src[10] + tmp[5] * src[8];
    
    det = src[0] * dst[0] + src[1] * dst[1] + src[2] * dst[2] + src[3] * dst[3]
    
    if det == 0 {
        panic("singular matrix is not invertible")
    }
    
    det = 1 / det
    
    for j := 0; j < 16;j++ {
        dst[j] *= det
    }
    
    return ret
}

func (this Mat4) AsArray() (out [16]float64) {
    // copy(out, this.Values)
    var i byte
    i = 0
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    out[i] = this.Values[i];i++
    
    return out
}

func (this Mat4) GetRotation() (r Mat4) {
    r.makeIdentity()
    r.Set(0,0, this.Get(0,0))
    r.Set(1,0, this.Get(1,0))
    r.Set(2,0, this.Get(2,0))
    r.Set(0,1, this.Get(0,1))
    r.Set(1,1, this.Get(1,1))
    r.Set(2,1, this.Get(2,1))
    r.Set(0,2, this.Get(0,2))
    r.Set(1,2, this.Get(1,2))
    r.Set(2,2, this.Get(2,2))
    return r
}

func (this Mat4) GetTranslation() (t Mat4) {
    t.makeIdentity()
    t.Set(3,0, this.Get(3,0))
    t.Set(3,1, this.Get(3,1))
    t.Set(3,2, this.Get(3,2))
    return t
}

func (this Mat4) GetPosition() vec3.Vec3 {
    return vec3.Vec3{
        this.Get(3,0),
        this.Get(3,1),
        this.Get(3,2),
    }
}

func (this Mat4) String() (s string) {
    s += "Matrix\n"
    for r := 0; r < 4; r++ {
        s += "( "
        for c := 0; c < 4;c++ {
            // precise
            // s += strconv.FormatFloat(this.Get(c,r), 'f', -1, 64) + " "
            // rounded
            s += strconv.FormatFloat(this.Get(c,r), 'f', 2, 64) + " "
        }
        s += ")\n"
    }
    return s
}

func (this Mat4) Equals( m Mat4, epsilon float64) bool {
    for i := 0;i != 16;i++ {
        if math.Abs(this.Values[i] - m.Values[i]) > epsilon {
            return false
        }
    }
    return true
}















