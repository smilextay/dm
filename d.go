/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1201 struct {
	dm_build_1202 *list.List
	dm_build_1203 *dm_build_1255
	dm_build_1204 int
}

func Dm_build_1205() *Dm_build_1201 {
	return &Dm_build_1201{
		dm_build_1202: list.New(),
		dm_build_1204: 0,
	}
}

func (dm_build_1207 *Dm_build_1201) Dm_build_1206() int {
	return dm_build_1207.dm_build_1204
}

func (dm_build_1209 *Dm_build_1201) Dm_build_1208(dm_build_1210 *Dm_build_1279, dm_build_1211 int) int {
	var dm_build_1212 = 0
	var dm_build_1213 = 0
	for dm_build_1212 < dm_build_1211 && dm_build_1209.dm_build_1203 != nil {
		dm_build_1213 = dm_build_1209.dm_build_1203.dm_build_1263(dm_build_1210, dm_build_1211-dm_build_1212)
		if dm_build_1209.dm_build_1203.dm_build_1258 == 0 {
			dm_build_1209.dm_build_1245()
		}
		dm_build_1212 += dm_build_1213
		dm_build_1209.dm_build_1204 -= dm_build_1213
	}
	return dm_build_1212
}

func (dm_build_1215 *Dm_build_1201) Dm_build_1214(dm_build_1216 []byte, dm_build_1217 int, dm_build_1218 int) int {
	var dm_build_1219 = 0
	var dm_build_1220 = 0
	for dm_build_1219 < dm_build_1218 && dm_build_1215.dm_build_1203 != nil {
		dm_build_1220 = dm_build_1215.dm_build_1203.dm_build_1267(dm_build_1216, dm_build_1217, dm_build_1218-dm_build_1219)
		if dm_build_1215.dm_build_1203.dm_build_1258 == 0 {
			dm_build_1215.dm_build_1245()
		}
		dm_build_1219 += dm_build_1220
		dm_build_1215.dm_build_1204 -= dm_build_1220
		dm_build_1217 += dm_build_1220
	}
	return dm_build_1219
}

func (dm_build_1222 *Dm_build_1201) Dm_build_1221(dm_build_1223 io.Writer, dm_build_1224 int) int {
	var dm_build_1225 = 0
	var dm_build_1226 = 0
	for dm_build_1225 < dm_build_1224 && dm_build_1222.dm_build_1203 != nil {
		dm_build_1226 = dm_build_1222.dm_build_1203.dm_build_1272(dm_build_1223, dm_build_1224-dm_build_1225)
		if dm_build_1222.dm_build_1203.dm_build_1258 == 0 {
			dm_build_1222.dm_build_1245()
		}
		dm_build_1225 += dm_build_1226
		dm_build_1222.dm_build_1204 -= dm_build_1226
	}
	return dm_build_1225
}

func (dm_build_1228 *Dm_build_1201) Dm_build_1227(dm_build_1229 []byte, dm_build_1230 int, dm_build_1231 int) {
	if dm_build_1231 == 0 {
		return
	}
	var dm_build_1232 = dm_build_1259(dm_build_1229, dm_build_1230, dm_build_1231)
	if dm_build_1228.dm_build_1203 == nil {
		dm_build_1228.dm_build_1203 = dm_build_1232
	} else {
		dm_build_1228.dm_build_1202.PushBack(dm_build_1232)
	}
	dm_build_1228.dm_build_1204 += dm_build_1231
}

func (dm_build_1234 *Dm_build_1201) dm_build_1233(dm_build_1235 int) byte {
	var dm_build_1236 = dm_build_1235
	var dm_build_1237 = dm_build_1234.dm_build_1203
	for dm_build_1236 > 0 && dm_build_1237 != nil {
		if dm_build_1237.dm_build_1258 == 0 {
			continue
		}
		if dm_build_1236 > dm_build_1237.dm_build_1258-1 {
			dm_build_1236 -= dm_build_1237.dm_build_1258
			dm_build_1237 = dm_build_1234.dm_build_1202.Front().Value.(*dm_build_1255)
		} else {
			break
		}
	}
	return dm_build_1237.dm_build_1276(dm_build_1236)
}
func (dm_build_1239 *Dm_build_1201) Dm_build_1238(dm_build_1240 *Dm_build_1201) {
	if dm_build_1240.dm_build_1204 == 0 {
		return
	}
	var dm_build_1241 = dm_build_1240.dm_build_1203
	for dm_build_1241 != nil {
		dm_build_1239.dm_build_1242(dm_build_1241)
		dm_build_1240.dm_build_1245()
		dm_build_1241 = dm_build_1240.dm_build_1203
	}
	dm_build_1240.dm_build_1204 = 0
}
func (dm_build_1243 *Dm_build_1201) dm_build_1242(dm_build_1244 *dm_build_1255) {
	if dm_build_1244.dm_build_1258 == 0 {
		return
	}
	if dm_build_1243.dm_build_1203 == nil {
		dm_build_1243.dm_build_1203 = dm_build_1244
	} else {
		dm_build_1243.dm_build_1202.PushBack(dm_build_1244)
	}
	dm_build_1243.dm_build_1204 += dm_build_1244.dm_build_1258
}

func (dm_build_1246 *Dm_build_1201) dm_build_1245() {
	var dm_build_1247 = dm_build_1246.dm_build_1202.Front()
	if dm_build_1247 == nil {
		dm_build_1246.dm_build_1203 = nil
	} else {
		dm_build_1246.dm_build_1203 = dm_build_1247.Value.(*dm_build_1255)
		dm_build_1246.dm_build_1202.Remove(dm_build_1247)
	}
}

func (dm_build_1249 *Dm_build_1201) Dm_build_1248() []byte {
	var dm_build_1250 = make([]byte, dm_build_1249.dm_build_1204)
	var dm_build_1251 = dm_build_1249.dm_build_1203
	var dm_build_1252 = 0
	var dm_build_1253 = len(dm_build_1250)
	var dm_build_1254 = 0
	for dm_build_1251 != nil {
		if dm_build_1251.dm_build_1258 > 0 {
			if dm_build_1253 > dm_build_1251.dm_build_1258 {
				dm_build_1254 = dm_build_1251.dm_build_1258
			} else {
				dm_build_1254 = dm_build_1253
			}
			copy(dm_build_1250[dm_build_1252:dm_build_1252+dm_build_1254], dm_build_1251.dm_build_1256[dm_build_1251.dm_build_1257:dm_build_1251.dm_build_1257+dm_build_1254])
			dm_build_1252 += dm_build_1254
			dm_build_1253 -= dm_build_1254
		}
		if dm_build_1249.dm_build_1202.Front() == nil {
			dm_build_1251 = nil
		} else {
			dm_build_1251 = dm_build_1249.dm_build_1202.Front().Value.(*dm_build_1255)
		}
	}
	return dm_build_1250
}

type dm_build_1255 struct {
	dm_build_1256 []byte
	dm_build_1257 int
	dm_build_1258 int
}

func dm_build_1259(dm_build_1260 []byte, dm_build_1261 int, dm_build_1262 int) *dm_build_1255 {
	return &dm_build_1255{
		dm_build_1260,
		dm_build_1261,
		dm_build_1262,
	}
}

func (dm_build_1264 *dm_build_1255) dm_build_1263(dm_build_1265 *Dm_build_1279, dm_build_1266 int) int {
	if dm_build_1264.dm_build_1258 <= dm_build_1266 {
		dm_build_1266 = dm_build_1264.dm_build_1258
	}
	dm_build_1265.Dm_build_1362(dm_build_1264.dm_build_1256[dm_build_1264.dm_build_1257 : dm_build_1264.dm_build_1257+dm_build_1266])
	dm_build_1264.dm_build_1257 += dm_build_1266
	dm_build_1264.dm_build_1258 -= dm_build_1266
	return dm_build_1266
}

func (dm_build_1268 *dm_build_1255) dm_build_1267(dm_build_1269 []byte, dm_build_1270 int, dm_build_1271 int) int {
	if dm_build_1268.dm_build_1258 <= dm_build_1271 {
		dm_build_1271 = dm_build_1268.dm_build_1258
	}
	copy(dm_build_1269[dm_build_1270:dm_build_1270+dm_build_1271], dm_build_1268.dm_build_1256[dm_build_1268.dm_build_1257:dm_build_1268.dm_build_1257+dm_build_1271])
	dm_build_1268.dm_build_1257 += dm_build_1271
	dm_build_1268.dm_build_1258 -= dm_build_1271
	return dm_build_1271
}

func (dm_build_1273 *dm_build_1255) dm_build_1272(dm_build_1274 io.Writer, dm_build_1275 int) int {
	if dm_build_1273.dm_build_1258 <= dm_build_1275 {
		dm_build_1275 = dm_build_1273.dm_build_1258
	}
	dm_build_1274.Write(dm_build_1273.dm_build_1256[dm_build_1273.dm_build_1257 : dm_build_1273.dm_build_1257+dm_build_1275])
	dm_build_1273.dm_build_1257 += dm_build_1275
	dm_build_1273.dm_build_1258 -= dm_build_1275
	return dm_build_1275
}
func (dm_build_1277 *dm_build_1255) dm_build_1276(dm_build_1278 int) byte {
	return dm_build_1277.dm_build_1256[dm_build_1277.dm_build_1257+dm_build_1278]
}
