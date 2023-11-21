/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_918 struct{}

var Dm_build_919 = &dm_build_918{}

func (Dm_build_921 *dm_build_918) Dm_build_920(dm_build_922 []byte, dm_build_923 int, dm_build_924 byte) int {
	dm_build_922[dm_build_923] = dm_build_924
	return 1
}

func (Dm_build_926 *dm_build_918) Dm_build_925(dm_build_927 []byte, dm_build_928 int, dm_build_929 int8) int {
	dm_build_927[dm_build_928] = byte(dm_build_929)
	return 1
}

func (Dm_build_931 *dm_build_918) Dm_build_930(dm_build_932 []byte, dm_build_933 int, dm_build_934 int16) int {
	dm_build_932[dm_build_933] = byte(dm_build_934)
	dm_build_933++
	dm_build_932[dm_build_933] = byte(dm_build_934 >> 8)
	return 2
}

func (Dm_build_936 *dm_build_918) Dm_build_935(dm_build_937 []byte, dm_build_938 int, dm_build_939 int32) int {
	dm_build_937[dm_build_938] = byte(dm_build_939)
	dm_build_938++
	dm_build_937[dm_build_938] = byte(dm_build_939 >> 8)
	dm_build_938++
	dm_build_937[dm_build_938] = byte(dm_build_939 >> 16)
	dm_build_938++
	dm_build_937[dm_build_938] = byte(dm_build_939 >> 24)
	dm_build_938++
	return 4
}

func (Dm_build_941 *dm_build_918) Dm_build_940(dm_build_942 []byte, dm_build_943 int, dm_build_944 int64) int {
	dm_build_942[dm_build_943] = byte(dm_build_944)
	dm_build_943++
	dm_build_942[dm_build_943] = byte(dm_build_944 >> 8)
	dm_build_943++
	dm_build_942[dm_build_943] = byte(dm_build_944 >> 16)
	dm_build_943++
	dm_build_942[dm_build_943] = byte(dm_build_944 >> 24)
	dm_build_943++
	dm_build_942[dm_build_943] = byte(dm_build_944 >> 32)
	dm_build_943++
	dm_build_942[dm_build_943] = byte(dm_build_944 >> 40)
	dm_build_943++
	dm_build_942[dm_build_943] = byte(dm_build_944 >> 48)
	dm_build_943++
	dm_build_942[dm_build_943] = byte(dm_build_944 >> 56)
	return 8
}

func (Dm_build_946 *dm_build_918) Dm_build_945(dm_build_947 []byte, dm_build_948 int, dm_build_949 float32) int {
	return Dm_build_946.Dm_build_965(dm_build_947, dm_build_948, math.Float32bits(dm_build_949))
}

func (Dm_build_951 *dm_build_918) Dm_build_950(dm_build_952 []byte, dm_build_953 int, dm_build_954 float64) int {
	return Dm_build_951.Dm_build_970(dm_build_952, dm_build_953, math.Float64bits(dm_build_954))
}

func (Dm_build_956 *dm_build_918) Dm_build_955(dm_build_957 []byte, dm_build_958 int, dm_build_959 uint8) int {
	dm_build_957[dm_build_958] = byte(dm_build_959)
	return 1
}

func (Dm_build_961 *dm_build_918) Dm_build_960(dm_build_962 []byte, dm_build_963 int, dm_build_964 uint16) int {
	dm_build_962[dm_build_963] = byte(dm_build_964)
	dm_build_963++
	dm_build_962[dm_build_963] = byte(dm_build_964 >> 8)
	return 2
}

func (Dm_build_966 *dm_build_918) Dm_build_965(dm_build_967 []byte, dm_build_968 int, dm_build_969 uint32) int {
	dm_build_967[dm_build_968] = byte(dm_build_969)
	dm_build_968++
	dm_build_967[dm_build_968] = byte(dm_build_969 >> 8)
	dm_build_968++
	dm_build_967[dm_build_968] = byte(dm_build_969 >> 16)
	dm_build_968++
	dm_build_967[dm_build_968] = byte(dm_build_969 >> 24)
	return 3
}

func (Dm_build_971 *dm_build_918) Dm_build_970(dm_build_972 []byte, dm_build_973 int, dm_build_974 uint64) int {
	dm_build_972[dm_build_973] = byte(dm_build_974)
	dm_build_973++
	dm_build_972[dm_build_973] = byte(dm_build_974 >> 8)
	dm_build_973++
	dm_build_972[dm_build_973] = byte(dm_build_974 >> 16)
	dm_build_973++
	dm_build_972[dm_build_973] = byte(dm_build_974 >> 24)
	dm_build_973++
	dm_build_972[dm_build_973] = byte(dm_build_974 >> 32)
	dm_build_973++
	dm_build_972[dm_build_973] = byte(dm_build_974 >> 40)
	dm_build_973++
	dm_build_972[dm_build_973] = byte(dm_build_974 >> 48)
	dm_build_973++
	dm_build_972[dm_build_973] = byte(dm_build_974 >> 56)
	return 3
}

func (Dm_build_976 *dm_build_918) Dm_build_975(dm_build_977 []byte, dm_build_978 int, dm_build_979 []byte, dm_build_980 int, dm_build_981 int) int {
	copy(dm_build_977[dm_build_978:dm_build_978+dm_build_981], dm_build_979[dm_build_980:dm_build_980+dm_build_981])
	return dm_build_981
}

func (Dm_build_983 *dm_build_918) Dm_build_982(dm_build_984 []byte, dm_build_985 int, dm_build_986 []byte, dm_build_987 int, dm_build_988 int) int {
	dm_build_985 += Dm_build_983.Dm_build_965(dm_build_984, dm_build_985, uint32(dm_build_988))
	return 4 + Dm_build_983.Dm_build_975(dm_build_984, dm_build_985, dm_build_986, dm_build_987, dm_build_988)
}

func (Dm_build_990 *dm_build_918) Dm_build_989(dm_build_991 []byte, dm_build_992 int, dm_build_993 []byte, dm_build_994 int, dm_build_995 int) int {
	dm_build_992 += Dm_build_990.Dm_build_960(dm_build_991, dm_build_992, uint16(dm_build_995))
	return 2 + Dm_build_990.Dm_build_975(dm_build_991, dm_build_992, dm_build_993, dm_build_994, dm_build_995)
}

func (Dm_build_997 *dm_build_918) Dm_build_996(dm_build_998 []byte, dm_build_999 int, dm_build_1000 string, dm_build_1001 string, dm_build_1002 *DmConnection) int {
	dm_build_1003 := Dm_build_997.Dm_build_1135(dm_build_1000, dm_build_1001, dm_build_1002)
	dm_build_999 += Dm_build_997.Dm_build_965(dm_build_998, dm_build_999, uint32(len(dm_build_1003)))
	return 4 + Dm_build_997.Dm_build_975(dm_build_998, dm_build_999, dm_build_1003, 0, len(dm_build_1003))
}

func (Dm_build_1005 *dm_build_918) Dm_build_1004(dm_build_1006 []byte, dm_build_1007 int, dm_build_1008 string, dm_build_1009 string, dm_build_1010 *DmConnection) int {
	dm_build_1011 := Dm_build_1005.Dm_build_1135(dm_build_1008, dm_build_1009, dm_build_1010)

	dm_build_1007 += Dm_build_1005.Dm_build_960(dm_build_1006, dm_build_1007, uint16(len(dm_build_1011)))
	return 2 + Dm_build_1005.Dm_build_975(dm_build_1006, dm_build_1007, dm_build_1011, 0, len(dm_build_1011))
}

func (Dm_build_1013 *dm_build_918) Dm_build_1012(dm_build_1014 []byte, dm_build_1015 int) byte {
	return dm_build_1014[dm_build_1015]
}

func (Dm_build_1017 *dm_build_918) Dm_build_1016(dm_build_1018 []byte, dm_build_1019 int) int16 {
	var dm_build_1020 int16
	dm_build_1020 = int16(dm_build_1018[dm_build_1019] & 0xff)
	dm_build_1019++
	dm_build_1020 |= int16(dm_build_1018[dm_build_1019]&0xff) << 8
	return dm_build_1020
}

func (Dm_build_1022 *dm_build_918) Dm_build_1021(dm_build_1023 []byte, dm_build_1024 int) int32 {
	var dm_build_1025 int32
	dm_build_1025 = int32(dm_build_1023[dm_build_1024] & 0xff)
	dm_build_1024++
	dm_build_1025 |= int32(dm_build_1023[dm_build_1024]&0xff) << 8
	dm_build_1024++
	dm_build_1025 |= int32(dm_build_1023[dm_build_1024]&0xff) << 16
	dm_build_1024++
	dm_build_1025 |= int32(dm_build_1023[dm_build_1024]&0xff) << 24
	return dm_build_1025
}

func (Dm_build_1027 *dm_build_918) Dm_build_1026(dm_build_1028 []byte, dm_build_1029 int) int64 {
	var dm_build_1030 int64
	dm_build_1030 = int64(dm_build_1028[dm_build_1029] & 0xff)
	dm_build_1029++
	dm_build_1030 |= int64(dm_build_1028[dm_build_1029]&0xff) << 8
	dm_build_1029++
	dm_build_1030 |= int64(dm_build_1028[dm_build_1029]&0xff) << 16
	dm_build_1029++
	dm_build_1030 |= int64(dm_build_1028[dm_build_1029]&0xff) << 24
	dm_build_1029++
	dm_build_1030 |= int64(dm_build_1028[dm_build_1029]&0xff) << 32
	dm_build_1029++
	dm_build_1030 |= int64(dm_build_1028[dm_build_1029]&0xff) << 40
	dm_build_1029++
	dm_build_1030 |= int64(dm_build_1028[dm_build_1029]&0xff) << 48
	dm_build_1029++
	dm_build_1030 |= int64(dm_build_1028[dm_build_1029]&0xff) << 56
	return dm_build_1030
}

func (Dm_build_1032 *dm_build_918) Dm_build_1031(dm_build_1033 []byte, dm_build_1034 int) float32 {
	return math.Float32frombits(Dm_build_1032.Dm_build_1048(dm_build_1033, dm_build_1034))
}

func (Dm_build_1036 *dm_build_918) Dm_build_1035(dm_build_1037 []byte, dm_build_1038 int) float64 {
	return math.Float64frombits(Dm_build_1036.Dm_build_1053(dm_build_1037, dm_build_1038))
}

func (Dm_build_1040 *dm_build_918) Dm_build_1039(dm_build_1041 []byte, dm_build_1042 int) uint8 {
	return uint8(dm_build_1041[dm_build_1042] & 0xff)
}

func (Dm_build_1044 *dm_build_918) Dm_build_1043(dm_build_1045 []byte, dm_build_1046 int) uint16 {
	var dm_build_1047 uint16
	dm_build_1047 = uint16(dm_build_1045[dm_build_1046] & 0xff)
	dm_build_1046++
	dm_build_1047 |= uint16(dm_build_1045[dm_build_1046]&0xff) << 8
	return dm_build_1047
}

func (Dm_build_1049 *dm_build_918) Dm_build_1048(dm_build_1050 []byte, dm_build_1051 int) uint32 {
	var dm_build_1052 uint32
	dm_build_1052 = uint32(dm_build_1050[dm_build_1051] & 0xff)
	dm_build_1051++
	dm_build_1052 |= uint32(dm_build_1050[dm_build_1051]&0xff) << 8
	dm_build_1051++
	dm_build_1052 |= uint32(dm_build_1050[dm_build_1051]&0xff) << 16
	dm_build_1051++
	dm_build_1052 |= uint32(dm_build_1050[dm_build_1051]&0xff) << 24
	return dm_build_1052
}

func (Dm_build_1054 *dm_build_918) Dm_build_1053(dm_build_1055 []byte, dm_build_1056 int) uint64 {
	var dm_build_1057 uint64
	dm_build_1057 = uint64(dm_build_1055[dm_build_1056] & 0xff)
	dm_build_1056++
	dm_build_1057 |= uint64(dm_build_1055[dm_build_1056]&0xff) << 8
	dm_build_1056++
	dm_build_1057 |= uint64(dm_build_1055[dm_build_1056]&0xff) << 16
	dm_build_1056++
	dm_build_1057 |= uint64(dm_build_1055[dm_build_1056]&0xff) << 24
	dm_build_1056++
	dm_build_1057 |= uint64(dm_build_1055[dm_build_1056]&0xff) << 32
	dm_build_1056++
	dm_build_1057 |= uint64(dm_build_1055[dm_build_1056]&0xff) << 40
	dm_build_1056++
	dm_build_1057 |= uint64(dm_build_1055[dm_build_1056]&0xff) << 48
	dm_build_1056++
	dm_build_1057 |= uint64(dm_build_1055[dm_build_1056]&0xff) << 56
	return dm_build_1057
}

func (Dm_build_1059 *dm_build_918) Dm_build_1058(dm_build_1060 []byte, dm_build_1061 int) []byte {
	dm_build_1062 := Dm_build_1059.Dm_build_1048(dm_build_1060, dm_build_1061)

	dm_build_1063 := make([]byte, dm_build_1062)
	copy(dm_build_1063[:int(dm_build_1062)], dm_build_1060[dm_build_1061+4:dm_build_1061+4+int(dm_build_1062)])
	return dm_build_1063
}

func (Dm_build_1065 *dm_build_918) Dm_build_1064(dm_build_1066 []byte, dm_build_1067 int) []byte {
	dm_build_1068 := Dm_build_1065.Dm_build_1043(dm_build_1066, dm_build_1067)

	dm_build_1069 := make([]byte, dm_build_1068)
	copy(dm_build_1069[:int(dm_build_1068)], dm_build_1066[dm_build_1067+2:dm_build_1067+2+int(dm_build_1068)])
	return dm_build_1069
}

func (Dm_build_1071 *dm_build_918) Dm_build_1070(dm_build_1072 []byte, dm_build_1073 int, dm_build_1074 int) []byte {

	dm_build_1075 := make([]byte, dm_build_1074)
	copy(dm_build_1075[:dm_build_1074], dm_build_1072[dm_build_1073:dm_build_1073+dm_build_1074])
	return dm_build_1075
}

func (Dm_build_1077 *dm_build_918) Dm_build_1076(dm_build_1078 []byte, dm_build_1079 int, dm_build_1080 int, dm_build_1081 string, dm_build_1082 *DmConnection) string {
	return Dm_build_1077.Dm_build_1172(dm_build_1078[dm_build_1079:dm_build_1079+dm_build_1080], dm_build_1081, dm_build_1082)
}

func (Dm_build_1084 *dm_build_918) Dm_build_1083(dm_build_1085 []byte, dm_build_1086 int, dm_build_1087 string, dm_build_1088 *DmConnection) string {
	dm_build_1089 := Dm_build_1084.Dm_build_1048(dm_build_1085, dm_build_1086)
	dm_build_1086 += 4
	return Dm_build_1084.Dm_build_1076(dm_build_1085, dm_build_1086, int(dm_build_1089), dm_build_1087, dm_build_1088)
}

func (Dm_build_1091 *dm_build_918) Dm_build_1090(dm_build_1092 []byte, dm_build_1093 int, dm_build_1094 string, dm_build_1095 *DmConnection) string {
	dm_build_1096 := Dm_build_1091.Dm_build_1043(dm_build_1092, dm_build_1093)
	dm_build_1093 += 2
	return Dm_build_1091.Dm_build_1076(dm_build_1092, dm_build_1093, int(dm_build_1096), dm_build_1094, dm_build_1095)
}

func (Dm_build_1098 *dm_build_918) Dm_build_1097(dm_build_1099 byte) []byte {
	return []byte{dm_build_1099}
}

func (Dm_build_1101 *dm_build_918) Dm_build_1100(dm_build_1102 int8) []byte {
	return []byte{byte(dm_build_1102)}
}

func (Dm_build_1104 *dm_build_918) Dm_build_1103(dm_build_1105 int16) []byte {
	return []byte{byte(dm_build_1105), byte(dm_build_1105 >> 8)}
}

func (Dm_build_1107 *dm_build_918) Dm_build_1106(dm_build_1108 int32) []byte {
	return []byte{byte(dm_build_1108), byte(dm_build_1108 >> 8), byte(dm_build_1108 >> 16), byte(dm_build_1108 >> 24)}
}

func (Dm_build_1110 *dm_build_918) Dm_build_1109(dm_build_1111 int64) []byte {
	return []byte{byte(dm_build_1111), byte(dm_build_1111 >> 8), byte(dm_build_1111 >> 16), byte(dm_build_1111 >> 24), byte(dm_build_1111 >> 32),
		byte(dm_build_1111 >> 40), byte(dm_build_1111 >> 48), byte(dm_build_1111 >> 56)}
}

func (Dm_build_1113 *dm_build_918) Dm_build_1112(dm_build_1114 float32) []byte {
	return Dm_build_1113.Dm_build_1124(math.Float32bits(dm_build_1114))
}

func (Dm_build_1116 *dm_build_918) Dm_build_1115(dm_build_1117 float64) []byte {
	return Dm_build_1116.Dm_build_1127(math.Float64bits(dm_build_1117))
}

func (Dm_build_1119 *dm_build_918) Dm_build_1118(dm_build_1120 uint8) []byte {
	return []byte{byte(dm_build_1120)}
}

func (Dm_build_1122 *dm_build_918) Dm_build_1121(dm_build_1123 uint16) []byte {
	return []byte{byte(dm_build_1123), byte(dm_build_1123 >> 8)}
}

func (Dm_build_1125 *dm_build_918) Dm_build_1124(dm_build_1126 uint32) []byte {
	return []byte{byte(dm_build_1126), byte(dm_build_1126 >> 8), byte(dm_build_1126 >> 16), byte(dm_build_1126 >> 24)}
}

func (Dm_build_1128 *dm_build_918) Dm_build_1127(dm_build_1129 uint64) []byte {
	return []byte{byte(dm_build_1129), byte(dm_build_1129 >> 8), byte(dm_build_1129 >> 16), byte(dm_build_1129 >> 24), byte(dm_build_1129 >> 32), byte(dm_build_1129 >> 40), byte(dm_build_1129 >> 48), byte(dm_build_1129 >> 56)}
}

func (Dm_build_1131 *dm_build_918) Dm_build_1130(dm_build_1132 []byte, dm_build_1133 string, dm_build_1134 *DmConnection) []byte {
	if dm_build_1133 == "UTF-8" {
		return dm_build_1132
	}

	if dm_build_1134 == nil {
		if e := dm_build_1177(dm_build_1133); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1132), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1134.encodeBuffer == nil {
		dm_build_1134.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1134.encode = dm_build_1177(dm_build_1134.getServerEncoding())
		dm_build_1134.transformReaderDst = make([]byte, 4096)
		dm_build_1134.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1134.encode; e != nil {

		dm_build_1134.encodeBuffer.Reset()

		n, err := dm_build_1134.encodeBuffer.ReadFrom(
			Dm_build_1191(bytes.NewReader(dm_build_1132), e.NewEncoder(), dm_build_1134.transformReaderDst, dm_build_1134.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1134.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1136 *dm_build_918) Dm_build_1135(dm_build_1137 string, dm_build_1138 string, dm_build_1139 *DmConnection) []byte {
	return Dm_build_1136.Dm_build_1130([]byte(dm_build_1137), dm_build_1138, dm_build_1139)
}

func (Dm_build_1141 *dm_build_918) Dm_build_1140(dm_build_1142 []byte) byte {
	return Dm_build_1141.Dm_build_1012(dm_build_1142, 0)
}

func (Dm_build_1144 *dm_build_918) Dm_build_1143(dm_build_1145 []byte) int16 {
	return Dm_build_1144.Dm_build_1016(dm_build_1145, 0)
}

func (Dm_build_1147 *dm_build_918) Dm_build_1146(dm_build_1148 []byte) int32 {
	return Dm_build_1147.Dm_build_1021(dm_build_1148, 0)
}

func (Dm_build_1150 *dm_build_918) Dm_build_1149(dm_build_1151 []byte) int64 {
	return Dm_build_1150.Dm_build_1026(dm_build_1151, 0)
}

func (Dm_build_1153 *dm_build_918) Dm_build_1152(dm_build_1154 []byte) float32 {
	return Dm_build_1153.Dm_build_1031(dm_build_1154, 0)
}

func (Dm_build_1156 *dm_build_918) Dm_build_1155(dm_build_1157 []byte) float64 {
	return Dm_build_1156.Dm_build_1035(dm_build_1157, 0)
}

func (Dm_build_1159 *dm_build_918) Dm_build_1158(dm_build_1160 []byte) uint8 {
	return Dm_build_1159.Dm_build_1039(dm_build_1160, 0)
}

func (Dm_build_1162 *dm_build_918) Dm_build_1161(dm_build_1163 []byte) uint16 {
	return Dm_build_1162.Dm_build_1043(dm_build_1163, 0)
}

func (Dm_build_1165 *dm_build_918) Dm_build_1164(dm_build_1166 []byte) uint32 {
	return Dm_build_1165.Dm_build_1048(dm_build_1166, 0)
}

func (Dm_build_1168 *dm_build_918) Dm_build_1167(dm_build_1169 []byte, dm_build_1170 string, dm_build_1171 *DmConnection) []byte {
	if dm_build_1170 == "UTF-8" {
		return dm_build_1169
	}

	if dm_build_1171 == nil {
		if e := dm_build_1177(dm_build_1170); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1169), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1171.encodeBuffer == nil {
		dm_build_1171.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1171.encode = dm_build_1177(dm_build_1171.getServerEncoding())
		dm_build_1171.transformReaderDst = make([]byte, 4096)
		dm_build_1171.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1171.encode; e != nil {

		dm_build_1171.encodeBuffer.Reset()

		n, err := dm_build_1171.encodeBuffer.ReadFrom(
			Dm_build_1191(bytes.NewReader(dm_build_1169), e.NewDecoder(), dm_build_1171.transformReaderDst, dm_build_1171.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1171.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1173 *dm_build_918) Dm_build_1172(dm_build_1174 []byte, dm_build_1175 string, dm_build_1176 *DmConnection) string {
	return string(Dm_build_1173.Dm_build_1167(dm_build_1174, dm_build_1175, dm_build_1176))
}

func dm_build_1177(dm_build_1178 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1178); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1179 struct {
	dm_build_1180 io.Reader
	dm_build_1181 transform.Transformer
	dm_build_1182 error

	dm_build_1183                []byte
	dm_build_1184, dm_build_1185 int

	dm_build_1186                []byte
	dm_build_1187, dm_build_1188 int

	dm_build_1189 bool
}

const dm_build_1190 = 4096

func Dm_build_1191(dm_build_1192 io.Reader, dm_build_1193 transform.Transformer, dm_build_1194 []byte, dm_build_1195 []byte) *Dm_build_1179 {
	dm_build_1193.Reset()
	return &Dm_build_1179{
		dm_build_1180: dm_build_1192,
		dm_build_1181: dm_build_1193,
		dm_build_1183: dm_build_1194,
		dm_build_1186: dm_build_1195,
	}
}

func (dm_build_1197 *Dm_build_1179) Read(dm_build_1198 []byte) (int, error) {
	dm_build_1199, dm_build_1200 := 0, error(nil)
	for {

		if dm_build_1197.dm_build_1184 != dm_build_1197.dm_build_1185 {
			dm_build_1199 = copy(dm_build_1198, dm_build_1197.dm_build_1183[dm_build_1197.dm_build_1184:dm_build_1197.dm_build_1185])
			dm_build_1197.dm_build_1184 += dm_build_1199
			if dm_build_1197.dm_build_1184 == dm_build_1197.dm_build_1185 && dm_build_1197.dm_build_1189 {
				return dm_build_1199, dm_build_1197.dm_build_1182
			}
			return dm_build_1199, nil
		} else if dm_build_1197.dm_build_1189 {
			return 0, dm_build_1197.dm_build_1182
		}

		if dm_build_1197.dm_build_1187 != dm_build_1197.dm_build_1188 || dm_build_1197.dm_build_1182 != nil {
			dm_build_1197.dm_build_1184 = 0
			dm_build_1197.dm_build_1185, dm_build_1199, dm_build_1200 = dm_build_1197.dm_build_1181.Transform(dm_build_1197.dm_build_1183, dm_build_1197.dm_build_1186[dm_build_1197.dm_build_1187:dm_build_1197.dm_build_1188], dm_build_1197.dm_build_1182 == io.EOF)
			dm_build_1197.dm_build_1187 += dm_build_1199

			switch {
			case dm_build_1200 == nil:
				if dm_build_1197.dm_build_1187 != dm_build_1197.dm_build_1188 {
					dm_build_1197.dm_build_1182 = nil
				}

				dm_build_1197.dm_build_1189 = dm_build_1197.dm_build_1182 != nil
				continue
			case dm_build_1200 == transform.ErrShortDst && (dm_build_1197.dm_build_1185 != 0 || dm_build_1199 != 0):

				continue
			case dm_build_1200 == transform.ErrShortSrc && dm_build_1197.dm_build_1188-dm_build_1197.dm_build_1187 != len(dm_build_1197.dm_build_1186) && dm_build_1197.dm_build_1182 == nil:

			default:
				dm_build_1197.dm_build_1189 = true

				if dm_build_1197.dm_build_1182 == nil || dm_build_1197.dm_build_1182 == io.EOF {
					dm_build_1197.dm_build_1182 = dm_build_1200
				}
				continue
			}
		}

		if dm_build_1197.dm_build_1187 != 0 {
			dm_build_1197.dm_build_1187, dm_build_1197.dm_build_1188 = 0, copy(dm_build_1197.dm_build_1186, dm_build_1197.dm_build_1186[dm_build_1197.dm_build_1187:dm_build_1197.dm_build_1188])
		}
		dm_build_1199, dm_build_1197.dm_build_1182 = dm_build_1197.dm_build_1180.Read(dm_build_1197.dm_build_1186[dm_build_1197.dm_build_1188:])
		dm_build_1197.dm_build_1188 += dm_build_1199
	}
}
