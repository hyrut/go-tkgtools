/*
    TKGTOOLS stands for Tgpp Key Generator Tools(MILENAGE defined in 3GPP TS 35.205).
    It implemente f1 - f5, f1*, f5* functions defined in 3GPP TS 35.205/35.206.
    These functions are also known as MILENAGE Algorithm Set.
    Test data could be find in TS 35.208.
    Specification is here: https://www.3gpp.org/ftp/Specs/archive/35_series
*/
package tkgtools

var S [256]byte = [256]byte{
 99,124,119,123,242,107,111,197, 48,  1,103, 43,254,215,171,118,
202,130,201,125,250, 89, 71,240,173,212,162,175,156,164,114,192,
183,253,147, 38, 54, 63,247,204, 52,165,229,241,113,216, 49, 21,
  4,199, 35,195, 24,150,  5,154,  7, 18,128,226,235, 39,178,117,
  9,131, 44, 26, 27,110, 90,160, 82, 59,214,179, 41,227, 47,132,
 83,209,  0,237, 32,252,177, 91,106,203,190, 57, 74, 76, 88,207,
208,239,170,251, 67, 77, 51,133, 69,249,  2,127, 80, 60,159,168,
 81,163, 64,143,146,157, 56,245,188,182,218, 33, 16,255,243,210,
205, 12, 19,236, 95,151, 68, 23,196,167,126, 61,100, 93, 25,115,
 96,129, 79,220, 34, 42,144,136, 70,238,184, 20,222, 94, 11,219,
224, 50, 58, 10, 73,  6, 36, 92,194,211,172, 98,145,149,228,121,
231,200, 55,109,141,213, 78,169,108, 86,244,234,101,122,174,  8,
186,120, 37, 46, 28,166,180,198,232,221,116, 31, 75,189,139,138,
112, 62,181,102, 72,  3,246, 14, 97, 53, 87,185,134,193, 29,158,
225,248,152, 17,105,217,142,148,155, 30,135,233,206, 85, 40,223,
140,161,137, 13,191,230, 66,104, 65,153, 45, 15,176, 84,187, 22,
}

var Xtime[256]byte = [256]byte{
  0,  2,  4,  6,  8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94,
 96, 98,100,102,104,106,108,110,112,114,116,118,120,122,124,126,
128,130,132,134,136,138,140,142,144,146,148,150,152,154,156,158,
160,162,164,166,168,170,172,174,176,178,180,182,184,186,188,190,
192,194,196,198,200,202,204,206,208,210,212,214,216,218,220,222,
224,226,228,230,232,234,236,238,240,242,244,246,248,250,252,254,
 27, 25, 31, 29, 19, 17, 23, 21, 11,  9, 15, 13,  3,  1,  7,  5,
 59, 57, 63, 61, 51, 49, 55, 53, 43, 41, 47, 45, 35, 33, 39, 37,
 91, 89, 95, 93, 83, 81, 87, 85, 75, 73, 79, 77, 67, 65, 71, 69,
123,121,127,125,115,113,119,117,107,105,111,109, 99, 97,103,101,
155,153,159,157,147,145,151,149,139,137,143,141,131,129,135,133,
187,185,191,189,179,177,183,181,171,169,175,173,163,161,167,165,
219,217,223,221,211,209,215,213,203,201,207,205,195,193,199,197,
251,249,255,253,243,241,247,245,235,233,239,237,227,225,231,229,
}

/*
  "class" TKGTOOLS contains r1-r5, c1-c5 value which are used in f1-f5, f1*, f5* functions.
*/
type TKGTOOLS struct{
  R1 uint8
  R2 uint8
  R3 uint8
  R4 uint8
  R5 uint8
  C1 [16]byte
  C2 [16]byte
  C3 [16]byte
  C4 [16]byte
  C5 [16]byte
  roundKeys [11][4][4]byte
}

/*
  Use this function by creating a TKGTOOLS "object".
  And with TKGTOOLS "object", you can call F1, F2345, F1star and F5star functions.
  Additional, r1-r5, c1-c5 values used in functions is defined in "object".
  And default value of r1-r5, c1-c5 is below:
    inst.R1 = 64
    inst.R2 = 0
    inst.R3 = 32
    inst.R4 = 64
    inst.R5 = 96
    inst.C1 = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
    inst.C2 = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
    inst.C3 = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02}
    inst.C4 = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04}
    inst.C5 = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08}
  Modify it if necessary after NewTKGTOOLS function call.
  All functions need byte array pointer, not byte array copy, that make functions runs in a quick way.
*/
func NewTKGTOOLS() *TKGTOOLS{
  inst := new(TKGTOOLS)
  /*
     Default r1 - r5, c1 - c5 in 3gpp TS35.206,
     r1 = 64; r2 = 0; r3 = 32; r4 = 64; r5 = 96
     c1 = 0, c2 = 1, c3 = 2, c4 = 4, c4 = 8
  */
  inst.R1 = 64
  inst.R2 = 0
  inst.R3 = 32
  inst.R4 = 64
  inst.R5 = 96
  inst.C1 = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
  inst.C2 = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
  inst.C3 = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02}
  inst.C4 = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04}
  inst.C5 = [16]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08}
  return inst
}

func _ringShiftLeft128Bits(data *[16]byte, nbit uint8){
  var i uint8
  tmp := *data
  leftShiftBytes := nbit / 8
  leftShiftBits := nbit % 8
  for i=0; i<16; i++{
    (*data)[(i+(16-leftShiftBytes))%16] = tmp[i]
  }
  mostLeftByte := (*data)[0]
  for i=0; i<15;i++{
    (*data)[i] = (*data)[i]<<leftShiftBits
    t := (*data)[i+1]
    t >>= 8 - leftShiftBits
    (*data)[i] |= t
  }
  (*data)[15] = ((*data)[15])<<leftShiftBits
  mostLeftByte >>= 8 - leftShiftBits
  (*data)[15] |= mostLeftByte
  return
}

func (tp *TKGTOOLS)_rijndaelKeySchedule(key *[16]byte){
  for i:=0; i<16; i++{
    tp.roundKeys[0][i&0x03][i>>2] = key[i]
  }

  roundConst := byte(1)
  for i:=1; i<11; i++{
    tp.roundKeys[i][0][0] = S[tp.roundKeys[i-1][1][3]] ^ tp.roundKeys[i-1][0][0] ^ roundConst
		tp.roundKeys[i][1][0] = S[tp.roundKeys[i-1][2][3]] ^ tp.roundKeys[i-1][1][0]
		tp.roundKeys[i][2][0] = S[tp.roundKeys[i-1][3][3]] ^ tp.roundKeys[i-1][2][0]
		tp.roundKeys[i][3][0] = S[tp.roundKeys[i-1][0][3]] ^ tp.roundKeys[i-1][3][0]

		for j:=0; j<4; j++{
			tp.roundKeys[i][j][1] = tp.roundKeys[i-1][j][1] ^ tp.roundKeys[i][j][0]
			tp.roundKeys[i][j][2] = tp.roundKeys[i-1][j][2] ^ tp.roundKeys[i][j][1]
			tp.roundKeys[i][j][3] = tp.roundKeys[i-1][j][3] ^ tp.roundKeys[i][j][2]
		}
		roundConst = Xtime[roundConst]
  }
}

func _keyAdd(state *[4][4]byte, roundKeys *[11][4][4]byte, round int){
  for i:=0; i<4; i++{
    for j:=0; j<4; j++{
			state[i][j] ^= roundKeys[round][i][j]
    }
  }
}

func _byteSub(state *[4][4]byte) int{
  for i:=0; i<4; i++{
    for j:=0; j<4; j++{
      state[i][j] = S[state[i][j]]
    }
  }
  return 0
}

func _shiftRow(state *[4][4]byte){
  var temp byte
  /* left rotate row 1 by 1 */
  temp = state[1][0]
  state[1][0] = state[1][1]
  state[1][1] = state[1][2]
  state[1][2] = state[1][3]
  state[1][3] = temp
  /* left rotate row 2 by 2 */
  temp = state[2][0]
  state[2][0] = state[2][2]
  state[2][2] = temp
  temp = state[2][1]
  state[2][1] = state[2][3]
  state[2][3] = temp
  /* left rotate row 3 by 3 */
  temp = state[3][0]
  state[3][0] = state[3][3]
  state[3][3] = state[3][2]
  state[3][2] = state[3][1]
  state[3][1] = temp
  return
}

/* MixColumn transformation*/
func _mixColumn(state *[4][4]byte){
  var temp, tmp, tmp0 byte
  /* do one column at a time */
  for i:=0; i<4;i++{
    temp = state[0][i] ^ state[1][i] ^ state[2][i] ^ state[3][i]
    tmp0 = state[0][i]

    /* Xtime array does multiply by x in GF2^8 */
    tmp = Xtime[state[0][i] ^ state[1][i]]
    state[0][i] ^= temp ^ tmp

    tmp = Xtime[state[1][i] ^ state[2][i]]
    state[1][i] ^= temp ^ tmp

    tmp = Xtime[state[2][i] ^ state[3][i]]
    state[2][i] ^= temp ^ tmp

    tmp = Xtime[state[3][i] ^ tmp0]
    state[3][i] ^= temp ^ tmp
  }
  return
}

func (tp *TKGTOOLS)_rijndaelEncrypt(input *[16]byte, output *[16]byte){
  var state [4][4]byte
  /* initialise state array from input byte string */
  for i:=0; i<16; i++{
    state[i & 0x3][i>>2] = input[i]
  }
  /* add first round_key */
  _keyAdd(&state, &tp.roundKeys, 0);
  /* do lots of full rounds */
  r := 1
  for r=1; r<=9; r++{
    _byteSub(&state)
    _shiftRow(&state)
    _mixColumn(&state)
    _keyAdd(&state, &tp.roundKeys, r)
  }
  /* final round */
  _byteSub(&state);
  _shiftRow(&state);
  _keyAdd(&state, &tp.roundKeys, r)
  /* produce output byte string from state array */
  for i:=0; i<16; i++{
    output[i] = state[i & 0x3][i>>2]
  }
  return
}

func (tp *TKGTOOLS)_computeOPc(op *[16]byte, op_c *[16]byte){
  tp._rijndaelEncrypt(op, op_c)
  for i:=0; i<16; i++{
    op_c[i] ^= op[i]
  }
  return
}

/*
  Function F1 is used for mac_a calculation.
  Input is key, rand, sqn, amf, op / opc
  Call it like:
    tkg = tkgtools.NewTKGTOOLS()
    tkg.F1(&key, &rand, &sqn , &amf, &mac_a, &op, nil)
    // tkg.F1(&key, &rand, &sqn , &amf, &mac_a, nil, &opc)
  Transfer op if you have, or opc if you have, keep another one as nil.
*/
func (tp *TKGTOOLS)F1(key *[16]byte, rand *[16]byte, sqn *[6]byte, amf *[2]byte, mac_a *[8]byte, op *[16]byte, opc *[16]byte){
  var op_c [16]byte
  var temp [16]byte
  var in1 [16]byte
  var out1 [16]byte
  var rijndaelInput [16]byte
  var i uint8

  tp._rijndaelKeySchedule(key)
  if opc==nil{
    tp._computeOPc(op, &op_c)
  }else{
    op_c = *opc
  }
  for i=0; i<16; i++{
    rijndaelInput[i] = rand[i] ^ op_c[i]
  }
  tp._rijndaelEncrypt(&rijndaelInput, &temp)
  for i=0; i<6; i++{
    in1[i] = sqn[i]
    in1[i+8] = sqn[i]
  }
  for i=0; i<2; i++{
    in1[i+6] = amf[i]
    in1[i+14] = amf[i]
  }
  /* XOR op_c and in1, rotate by r1=64, and XOR *
   * on the constant c1 (which is all zeroes)   */
  for i=0; i<16; i++{
    rijndaelInput[i] = in1[i] ^ op_c[i]
  }
  _ringShiftLeft128Bits(&rijndaelInput, tp.R1)
  for i=0; i<16; i++{
    rijndaelInput[i] ^= tp.C1[i]
  }

  /* XOR on the value temp computed before */
  for i=0; i<16; i++{
    rijndaelInput[i] ^= temp[i]
  }
  tp._rijndaelEncrypt(&rijndaelInput, &out1)
  for i=0; i<16; i++{
    out1[i] ^= op_c[i]
  }
  for i=0; i<8; i++{
    mac_a[i] = out1[i]
  }
  return
}

/*
  Function F2345 is used for res, ck, ik and ak calculation.
  Input is key, rand, op / opc
  Call it like:
    tkg = tkgtools.NewTKGTOOLS()
    tkg.F2345(&key, &rand, &res, &ck, &ik, &ak, &op, nil)
    // tkg.F2345(&key, &rand, &res, &ck, &ik, &ak, nil, &opc)
  Transfer op if you have, or opc if you have, keep another one as nil.
*/
func (tp *TKGTOOLS)F2345 (key *[16]byte, rand *[16]byte, res *[8]byte, ck *[16]byte, ik *[16]byte, ak *[6]byte, op *[16]byte, opc *[16]byte){
  var op_c [16]byte
  var temp [16]byte
  var out [16]byte
  var rijndaelInput [16]byte
  var i uint8

  tp._rijndaelKeySchedule(key)
  if opc==nil{
    tp._computeOPc(op, &op_c)
  }else{
    op_c = *opc
  }
  for i=0; i<16; i++{
    rijndaelInput[i] = rand[i] ^ op_c[i]
  }
  tp._rijndaelEncrypt(&rijndaelInput, &temp);
  /* To obtain output block OUT2: XOR OPc and TEMP,    *
   * rotate by r2=0, and XOR on the constant c2 (which *
   * is all zeroes except that the last bit is 1).     */
  for i=0; i<16; i++{
    rijndaelInput[i] = temp[i] ^ op_c[i]
  }
  _ringShiftLeft128Bits(&rijndaelInput, tp.R2)
  for i=0; i<16; i++{
    rijndaelInput[i] ^= tp.C2[i]
  }
  tp._rijndaelEncrypt(&rijndaelInput, &out)
  for i=0; i<16; i++{
    out[i] ^= op_c[i]
  }
  for i=0; i<8; i++{
    res[i] = out[i+8]
  }
  for i=0; i<6; i++{
    ak[i]  = out[i]
  }
  /* To obtain output block OUT3: XOR OPc and TEMP,        *
   * rotate by r3=32, and XOR on the constant c3 (which    *
   * is all zeroes except that the next to last bit is 1). */
  for i=0; i<16; i++{
    rijndaelInput[i] = temp[i] ^ op_c[i]
  }
  _ringShiftLeft128Bits(&rijndaelInput, tp.R3)
  for i=0; i<16; i++{
    rijndaelInput[i] ^= tp.C3[i]
  }
  tp._rijndaelEncrypt(&rijndaelInput, &out)
  for i=0; i<16; i++{
    out[i] ^= op_c[i]
  }
  for i=0; i<16; i++{
    ck[i] = out[i]
  }
  /* To obtain output block OUT4: XOR OPc and TEMP,         *
   * rotate by r4=64, and XOR on the constant c4 (which     *
   * is all zeroes except that the 2nd from last bit is 1). */
  for i=0; i<16; i++{
    rijndaelInput[i] = temp[i] ^ op_c[i]
  }
  _ringShiftLeft128Bits(&rijndaelInput, tp.R4)
  for i=0; i<16; i++{
    rijndaelInput[i] ^= tp.C4[i]
  }
  tp._rijndaelEncrypt(&rijndaelInput, &out)
  for i=0; i<16; i++{
    out[i] ^= op_c[i]
  }
  for i=0; i<16; i++{
    ik[i] = out[i]
  }
  return
}

/*
  Function F1 is used for mac_s calculation.
  Input is key, rand, sqn, amf, op / opc
  Call it like:
    tkg = tkgtools.NewTKGTOOLS()
    tkg.F1star(&key, &rand, &sqn , &amf, &mac_s, &op, nil)
    // tkg.F1star(&key, &rand, &sqn , &amf, &mac_s, nil, &opc)
  Transfer op if you have, or opc if you have, keep another one as nil.
*/
func (tp *TKGTOOLS)F1star(key *[16]byte, rand *[16]byte, sqn *[6]byte, amf *[2]byte, mac_s *[8]byte, op *[16]byte, opc *[16]byte){
  var op_c [16]byte
  var temp[16]byte
  var in1[16]byte
  var out1[16]byte
  var rijndaelInput[16]byte
  var i uint8

  tp._rijndaelKeySchedule(key);
  if opc==nil{
    tp._computeOPc(op, &op_c)
  }else{
    op_c = *opc
  }
  for i=0; i<16; i++{
    rijndaelInput[i] = rand[i] ^ op_c[i]
  }
  tp._rijndaelEncrypt(&rijndaelInput, &temp)
  for i=0; i<6; i++{
    in1[i] = sqn[i]
    in1[i+8] = sqn[i]
  }
  for i=0; i<2; i++{
    in1[i+6] = amf[i]
    in1[i+14] = amf[i]
  }
  /* XOR op_c and in1, rotate by r1=64, and XOR *
   * on the constant c1 (which is all zeroes)   */
  for i=0; i<16; i++{
    rijndaelInput[i] = in1[i] ^ op_c[i]
  }
  _ringShiftLeft128Bits(&rijndaelInput, tp.R1)
  for i=0; i<16; i++{
    rijndaelInput[i] ^= tp.C1[i]
  }
  /* XOR on the value temp computed before */
  for i=0; i<16; i++{
    rijndaelInput[i] ^= temp[i]
  }
  tp._rijndaelEncrypt(&rijndaelInput, &out1)
  for i=0; i<16; i++{
    out1[i] ^= op_c[i]
  }
  for i=0; i<8; i++{
    mac_s[i] = out1[i+8]
  }
  return
}

/*
  Function F5star is used for ak calculation.
  Input is key, rand, op / opc
  Call it like:
    tkg = tkgtools.NewTKGTOOLS()
    tkg.F5star(&key, &rand, &ak, &op, nil)
    // tkg.F5star(&key, &rand, &ak, nil, &opc)
  Transfer op if you have, or opc if you have, keep another one as nil.
*/
func (tp *TKGTOOLS)F5star(key *[16]byte, rand *[16]byte, ak *[6]byte, op *[16]byte, opc *[16]byte){
  var op_c [16]byte
  var temp [16]byte
  var out [16]byte
  var rijndaelInput [16]byte
  var i uint8

  tp._rijndaelKeySchedule(key)
  if opc==nil{
    tp._computeOPc(op, &op_c)
  }else{
    op_c = *opc
  }
  for i=0; i<16; i++{
    rijndaelInput[i] = rand[i] ^ op_c[i]
  }
  tp._rijndaelEncrypt(&rijndaelInput, &temp)
  /* To obtain output block OUT5: XOR OPc and TEMP,         *
   * rotate by r5=96, and XOR on the constant c5 (which     *
   * is all zeroes except that the 3rd from last bit is 1). */
  for i=0; i<16; i++{
    rijndaelInput[i] = temp[i] ^ op_c[i]
  }
  _ringShiftLeft128Bits(&rijndaelInput, tp.R5)
  for i=0; i<16; i++{
    rijndaelInput[i] ^= tp.C5[i]
  }

  tp._rijndaelEncrypt(&rijndaelInput, &out)
  for i=0; i<16; i++{
    out[i] ^= op_c[i]
  }
  for i=0; i<6; i++{
    ak[i] = out[i]
  }
  return
}
