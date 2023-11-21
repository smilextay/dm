/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_1279 struct {
	dm_build_1280 []byte
	dm_build_1281 int
}

func Dm_build_1282(dm_build_1283 int) *Dm_build_1279 {
	return &Dm_build_1279{make([]byte, 0, dm_build_1283), 0}
}

func Dm_build_1284(dm_build_1285 []byte) *Dm_build_1279 {
	return &Dm_build_1279{dm_build_1285, 0}
}

func (dm_build_1287 *Dm_build_1279) dm_build_1286(dm_build_1288 int) *Dm_build_1279 {

	dm_build_1289 := len(dm_build_1287.dm_build_1280)
	dm_build_1290 := cap(dm_build_1287.dm_build_1280)

	if dm_build_1289+dm_build_1288 <= dm_build_1290 {
		dm_build_1287.dm_build_1280 = dm_build_1287.dm_build_1280[:dm_build_1289+dm_build_1288]
	} else {

		var calCap = int64(math.Max(float64(2*dm_build_1290), float64(dm_build_1288+dm_build_1289)))

		nbuf := make([]byte, dm_build_1288+dm_build_1289, calCap)
		copy(nbuf, dm_build_1287.dm_build_1280)
		dm_build_1287.dm_build_1280 = nbuf
	}

	return dm_build_1287
}

func (dm_build_1292 *Dm_build_1279) Dm_build_1291() int {
	return len(dm_build_1292.dm_build_1280)
}

func (dm_build_1294 *Dm_build_1279) Dm_build_1293(dm_build_1295 int) *Dm_build_1279 {
	for i := dm_build_1295; i < len(dm_build_1294.dm_build_1280); i++ {
		dm_build_1294.dm_build_1280[i] = 0
	}
	dm_build_1294.dm_build_1280 = dm_build_1294.dm_build_1280[:dm_build_1295]
	return dm_build_1294
}

func (dm_build_1297 *Dm_build_1279) Dm_build_1296(dm_build_1298 int) *Dm_build_1279 {
	dm_build_1297.dm_build_1281 = dm_build_1298
	return dm_build_1297
}

func (dm_build_1300 *Dm_build_1279) Dm_build_1299() int {
	return dm_build_1300.dm_build_1281
}

func (dm_build_1302 *Dm_build_1279) Dm_build_1301(dm_build_1303 bool) int {
	return len(dm_build_1302.dm_build_1280) - dm_build_1302.dm_build_1281
}

func (dm_build_1305 *Dm_build_1279) Dm_build_1304(dm_build_1306 int, dm_build_1307 bool, dm_build_1308 bool) *Dm_build_1279 {

	if dm_build_1307 {
		if dm_build_1308 {
			dm_build_1305.dm_build_1286(dm_build_1306)
		} else {
			dm_build_1305.dm_build_1280 = dm_build_1305.dm_build_1280[:len(dm_build_1305.dm_build_1280)-dm_build_1306]
		}
	} else {
		if dm_build_1308 {
			dm_build_1305.dm_build_1281 += dm_build_1306
		} else {
			dm_build_1305.dm_build_1281 -= dm_build_1306
		}
	}

	return dm_build_1305
}

func (dm_build_1310 *Dm_build_1279) Dm_build_1309(dm_build_1311 io.Reader, dm_build_1312 int) (int, error) {
	dm_build_1313 := len(dm_build_1310.dm_build_1280)
	dm_build_1310.dm_build_1286(dm_build_1312)
	dm_build_1314 := 0
	for dm_build_1312 > 0 {
		n, err := dm_build_1311.Read(dm_build_1310.dm_build_1280[dm_build_1313+dm_build_1314:])
		if n > 0 && err == io.EOF {
			dm_build_1314 += n
			dm_build_1310.dm_build_1280 = dm_build_1310.dm_build_1280[:dm_build_1313+dm_build_1314]
			return dm_build_1314, nil
		} else if n > 0 && err == nil {
			dm_build_1312 -= n
			dm_build_1314 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_1314, nil
}

func (dm_build_1316 *Dm_build_1279) Dm_build_1315(dm_build_1317 io.Writer) (*Dm_build_1279, error) {
	if _, err := dm_build_1317.Write(dm_build_1316.dm_build_1280); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return dm_build_1316, nil
}

func (dm_build_1319 *Dm_build_1279) Dm_build_1318(dm_build_1320 bool) int {
	dm_build_1321 := len(dm_build_1319.dm_build_1280)
	dm_build_1319.dm_build_1286(1)

	if dm_build_1320 {
		return copy(dm_build_1319.dm_build_1280[dm_build_1321:], []byte{1})
	} else {
		return copy(dm_build_1319.dm_build_1280[dm_build_1321:], []byte{0})
	}
}

func (dm_build_1323 *Dm_build_1279) Dm_build_1322(dm_build_1324 byte) int {
	dm_build_1325 := len(dm_build_1323.dm_build_1280)
	dm_build_1323.dm_build_1286(1)

	return copy(dm_build_1323.dm_build_1280[dm_build_1325:], Dm_build_919.Dm_build_1097(dm_build_1324))
}

func (dm_build_1327 *Dm_build_1279) Dm_build_1326(dm_build_1328 int8) int {
	dm_build_1329 := len(dm_build_1327.dm_build_1280)
	dm_build_1327.dm_build_1286(1)

	return copy(dm_build_1327.dm_build_1280[dm_build_1329:], Dm_build_919.Dm_build_1100(dm_build_1328))
}

func (dm_build_1331 *Dm_build_1279) Dm_build_1330(dm_build_1332 int16) int {
	dm_build_1333 := len(dm_build_1331.dm_build_1280)
	dm_build_1331.dm_build_1286(2)

	return copy(dm_build_1331.dm_build_1280[dm_build_1333:], Dm_build_919.Dm_build_1103(dm_build_1332))
}

func (dm_build_1335 *Dm_build_1279) Dm_build_1334(dm_build_1336 int32) int {
	dm_build_1337 := len(dm_build_1335.dm_build_1280)
	dm_build_1335.dm_build_1286(4)

	return copy(dm_build_1335.dm_build_1280[dm_build_1337:], Dm_build_919.Dm_build_1106(dm_build_1336))
}

func (dm_build_1339 *Dm_build_1279) Dm_build_1338(dm_build_1340 uint8) int {
	dm_build_1341 := len(dm_build_1339.dm_build_1280)
	dm_build_1339.dm_build_1286(1)

	return copy(dm_build_1339.dm_build_1280[dm_build_1341:], Dm_build_919.Dm_build_1118(dm_build_1340))
}

func (dm_build_1343 *Dm_build_1279) Dm_build_1342(dm_build_1344 uint16) int {
	dm_build_1345 := len(dm_build_1343.dm_build_1280)
	dm_build_1343.dm_build_1286(2)

	return copy(dm_build_1343.dm_build_1280[dm_build_1345:], Dm_build_919.Dm_build_1121(dm_build_1344))
}

func (dm_build_1347 *Dm_build_1279) Dm_build_1346(dm_build_1348 uint32) int {
	dm_build_1349 := len(dm_build_1347.dm_build_1280)
	dm_build_1347.dm_build_1286(4)

	return copy(dm_build_1347.dm_build_1280[dm_build_1349:], Dm_build_919.Dm_build_1124(dm_build_1348))
}

func (dm_build_1351 *Dm_build_1279) Dm_build_1350(dm_build_1352 uint64) int {
	dm_build_1353 := len(dm_build_1351.dm_build_1280)
	dm_build_1351.dm_build_1286(8)

	return copy(dm_build_1351.dm_build_1280[dm_build_1353:], Dm_build_919.Dm_build_1127(dm_build_1352))
}

func (dm_build_1355 *Dm_build_1279) Dm_build_1354(dm_build_1356 float32) int {
	dm_build_1357 := len(dm_build_1355.dm_build_1280)
	dm_build_1355.dm_build_1286(4)

	return copy(dm_build_1355.dm_build_1280[dm_build_1357:], Dm_build_919.Dm_build_1124(math.Float32bits(dm_build_1356)))
}

func (dm_build_1359 *Dm_build_1279) Dm_build_1358(dm_build_1360 float64) int {
	dm_build_1361 := len(dm_build_1359.dm_build_1280)
	dm_build_1359.dm_build_1286(8)

	return copy(dm_build_1359.dm_build_1280[dm_build_1361:], Dm_build_919.Dm_build_1127(math.Float64bits(dm_build_1360)))
}

func (dm_build_1363 *Dm_build_1279) Dm_build_1362(dm_build_1364 []byte) int {
	dm_build_1365 := len(dm_build_1363.dm_build_1280)
	dm_build_1363.dm_build_1286(len(dm_build_1364))
	return copy(dm_build_1363.dm_build_1280[dm_build_1365:], dm_build_1364)
}

func (dm_build_1367 *Dm_build_1279) Dm_build_1366(dm_build_1368 []byte) int {
	return dm_build_1367.Dm_build_1334(int32(len(dm_build_1368))) + dm_build_1367.Dm_build_1362(dm_build_1368)
}

func (dm_build_1370 *Dm_build_1279) Dm_build_1369(dm_build_1371 []byte) int {
	return dm_build_1370.Dm_build_1338(uint8(len(dm_build_1371))) + dm_build_1370.Dm_build_1362(dm_build_1371)
}

func (dm_build_1373 *Dm_build_1279) Dm_build_1372(dm_build_1374 []byte) int {
	return dm_build_1373.Dm_build_1342(uint16(len(dm_build_1374))) + dm_build_1373.Dm_build_1362(dm_build_1374)
}

func (dm_build_1376 *Dm_build_1279) Dm_build_1375(dm_build_1377 []byte) int {
	return dm_build_1376.Dm_build_1362(dm_build_1377) + dm_build_1376.Dm_build_1322(0)
}

func (dm_build_1379 *Dm_build_1279) Dm_build_1378(dm_build_1380 string, dm_build_1381 string, dm_build_1382 *DmConnection) int {
	dm_build_1383 := Dm_build_919.Dm_build_1135(dm_build_1380, dm_build_1381, dm_build_1382)
	return dm_build_1379.Dm_build_1366(dm_build_1383)
}

func (dm_build_1385 *Dm_build_1279) Dm_build_1384(dm_build_1386 string, dm_build_1387 string, dm_build_1388 *DmConnection) int {
	dm_build_1389 := Dm_build_919.Dm_build_1135(dm_build_1386, dm_build_1387, dm_build_1388)
	return dm_build_1385.Dm_build_1369(dm_build_1389)
}

func (dm_build_1391 *Dm_build_1279) Dm_build_1390(dm_build_1392 string, dm_build_1393 string, dm_build_1394 *DmConnection) int {
	dm_build_1395 := Dm_build_919.Dm_build_1135(dm_build_1392, dm_build_1393, dm_build_1394)
	return dm_build_1391.Dm_build_1372(dm_build_1395)
}

func (dm_build_1397 *Dm_build_1279) Dm_build_1396(dm_build_1398 string, dm_build_1399 string, dm_build_1400 *DmConnection) int {
	dm_build_1401 := Dm_build_919.Dm_build_1135(dm_build_1398, dm_build_1399, dm_build_1400)
	return dm_build_1397.Dm_build_1375(dm_build_1401)
}

func (dm_build_1403 *Dm_build_1279) Dm_build_1402() byte {
	dm_build_1404 := Dm_build_919.Dm_build_1012(dm_build_1403.dm_build_1280, dm_build_1403.dm_build_1281)
	dm_build_1403.dm_build_1281++
	return dm_build_1404
}

func (dm_build_1406 *Dm_build_1279) Dm_build_1405() int16 {
	dm_build_1407 := Dm_build_919.Dm_build_1016(dm_build_1406.dm_build_1280, dm_build_1406.dm_build_1281)
	dm_build_1406.dm_build_1281 += 2
	return dm_build_1407
}

func (dm_build_1409 *Dm_build_1279) Dm_build_1408() int32 {
	dm_build_1410 := Dm_build_919.Dm_build_1021(dm_build_1409.dm_build_1280, dm_build_1409.dm_build_1281)
	dm_build_1409.dm_build_1281 += 4
	return dm_build_1410
}

func (dm_build_1412 *Dm_build_1279) Dm_build_1411() int64 {
	dm_build_1413 := Dm_build_919.Dm_build_1026(dm_build_1412.dm_build_1280, dm_build_1412.dm_build_1281)
	dm_build_1412.dm_build_1281 += 8
	return dm_build_1413
}

func (dm_build_1415 *Dm_build_1279) Dm_build_1414() float32 {
	dm_build_1416 := Dm_build_919.Dm_build_1031(dm_build_1415.dm_build_1280, dm_build_1415.dm_build_1281)
	dm_build_1415.dm_build_1281 += 4
	return dm_build_1416
}

func (dm_build_1418 *Dm_build_1279) Dm_build_1417() float64 {
	dm_build_1419 := Dm_build_919.Dm_build_1035(dm_build_1418.dm_build_1280, dm_build_1418.dm_build_1281)
	dm_build_1418.dm_build_1281 += 8
	return dm_build_1419
}

func (dm_build_1421 *Dm_build_1279) Dm_build_1420() uint8 {
	dm_build_1422 := Dm_build_919.Dm_build_1039(dm_build_1421.dm_build_1280, dm_build_1421.dm_build_1281)
	dm_build_1421.dm_build_1281 += 1
	return dm_build_1422
}

func (dm_build_1424 *Dm_build_1279) Dm_build_1423() uint16 {
	dm_build_1425 := Dm_build_919.Dm_build_1043(dm_build_1424.dm_build_1280, dm_build_1424.dm_build_1281)
	dm_build_1424.dm_build_1281 += 2
	return dm_build_1425
}

func (dm_build_1427 *Dm_build_1279) Dm_build_1426() uint32 {
	dm_build_1428 := Dm_build_919.Dm_build_1048(dm_build_1427.dm_build_1280, dm_build_1427.dm_build_1281)
	dm_build_1427.dm_build_1281 += 4
	return dm_build_1428
}

func (dm_build_1430 *Dm_build_1279) Dm_build_1429(dm_build_1431 int) []byte {
	dm_build_1432 := Dm_build_919.Dm_build_1070(dm_build_1430.dm_build_1280, dm_build_1430.dm_build_1281, dm_build_1431)
	dm_build_1430.dm_build_1281 += dm_build_1431
	return dm_build_1432
}

func (dm_build_1434 *Dm_build_1279) Dm_build_1433() []byte {
	return dm_build_1434.Dm_build_1429(int(dm_build_1434.Dm_build_1408()))
}

func (dm_build_1436 *Dm_build_1279) Dm_build_1435() []byte {
	return dm_build_1436.Dm_build_1429(int(dm_build_1436.Dm_build_1402()))
}

func (dm_build_1438 *Dm_build_1279) Dm_build_1437() []byte {
	return dm_build_1438.Dm_build_1429(int(dm_build_1438.Dm_build_1405()))
}

func (dm_build_1440 *Dm_build_1279) Dm_build_1439(dm_build_1441 int) []byte {
	return dm_build_1440.Dm_build_1429(dm_build_1441)
}

func (dm_build_1443 *Dm_build_1279) Dm_build_1442() []byte {
	dm_build_1444 := 0
	for dm_build_1443.Dm_build_1402() != 0 {
		dm_build_1444++
	}
	dm_build_1443.Dm_build_1304(dm_build_1444, false, false)
	return dm_build_1443.Dm_build_1429(dm_build_1444)
}

func (dm_build_1446 *Dm_build_1279) Dm_build_1445(dm_build_1447 int, dm_build_1448 string, dm_build_1449 *DmConnection) string {
	return Dm_build_919.Dm_build_1172(dm_build_1446.Dm_build_1429(dm_build_1447), dm_build_1448, dm_build_1449)
}

func (dm_build_1451 *Dm_build_1279) Dm_build_1450(dm_build_1452 string, dm_build_1453 *DmConnection) string {
	return Dm_build_919.Dm_build_1172(dm_build_1451.Dm_build_1433(), dm_build_1452, dm_build_1453)
}

func (dm_build_1455 *Dm_build_1279) Dm_build_1454(dm_build_1456 string, dm_build_1457 *DmConnection) string {
	return Dm_build_919.Dm_build_1172(dm_build_1455.Dm_build_1435(), dm_build_1456, dm_build_1457)
}

func (dm_build_1459 *Dm_build_1279) Dm_build_1458(dm_build_1460 string, dm_build_1461 *DmConnection) string {
	return Dm_build_919.Dm_build_1172(dm_build_1459.Dm_build_1437(), dm_build_1460, dm_build_1461)
}

func (dm_build_1463 *Dm_build_1279) Dm_build_1462(dm_build_1464 string, dm_build_1465 *DmConnection) string {
	return Dm_build_919.Dm_build_1172(dm_build_1463.Dm_build_1442(), dm_build_1464, dm_build_1465)
}

func (dm_build_1467 *Dm_build_1279) Dm_build_1466(dm_build_1468 int, dm_build_1469 byte) int {
	return dm_build_1467.Dm_build_1502(dm_build_1468, Dm_build_919.Dm_build_1097(dm_build_1469))
}

func (dm_build_1471 *Dm_build_1279) Dm_build_1470(dm_build_1472 int, dm_build_1473 int16) int {
	return dm_build_1471.Dm_build_1502(dm_build_1472, Dm_build_919.Dm_build_1103(dm_build_1473))
}

func (dm_build_1475 *Dm_build_1279) Dm_build_1474(dm_build_1476 int, dm_build_1477 int32) int {
	return dm_build_1475.Dm_build_1502(dm_build_1476, Dm_build_919.Dm_build_1106(dm_build_1477))
}

func (dm_build_1479 *Dm_build_1279) Dm_build_1478(dm_build_1480 int, dm_build_1481 int64) int {
	return dm_build_1479.Dm_build_1502(dm_build_1480, Dm_build_919.Dm_build_1109(dm_build_1481))
}

func (dm_build_1483 *Dm_build_1279) Dm_build_1482(dm_build_1484 int, dm_build_1485 float32) int {
	return dm_build_1483.Dm_build_1502(dm_build_1484, Dm_build_919.Dm_build_1112(dm_build_1485))
}

func (dm_build_1487 *Dm_build_1279) Dm_build_1486(dm_build_1488 int, dm_build_1489 float64) int {
	return dm_build_1487.Dm_build_1502(dm_build_1488, Dm_build_919.Dm_build_1115(dm_build_1489))
}

func (dm_build_1491 *Dm_build_1279) Dm_build_1490(dm_build_1492 int, dm_build_1493 uint8) int {
	return dm_build_1491.Dm_build_1502(dm_build_1492, Dm_build_919.Dm_build_1118(dm_build_1493))
}

func (dm_build_1495 *Dm_build_1279) Dm_build_1494(dm_build_1496 int, dm_build_1497 uint16) int {
	return dm_build_1495.Dm_build_1502(dm_build_1496, Dm_build_919.Dm_build_1121(dm_build_1497))
}

func (dm_build_1499 *Dm_build_1279) Dm_build_1498(dm_build_1500 int, dm_build_1501 uint32) int {
	return dm_build_1499.Dm_build_1502(dm_build_1500, Dm_build_919.Dm_build_1124(dm_build_1501))
}

func (dm_build_1503 *Dm_build_1279) Dm_build_1502(dm_build_1504 int, dm_build_1505 []byte) int {
	return copy(dm_build_1503.dm_build_1280[dm_build_1504:], dm_build_1505)
}

func (dm_build_1507 *Dm_build_1279) Dm_build_1506(dm_build_1508 int, dm_build_1509 []byte) int {
	return dm_build_1507.Dm_build_1474(dm_build_1508, int32(len(dm_build_1509))) + dm_build_1507.Dm_build_1502(dm_build_1508+4, dm_build_1509)
}

func (dm_build_1511 *Dm_build_1279) Dm_build_1510(dm_build_1512 int, dm_build_1513 []byte) int {
	return dm_build_1511.Dm_build_1466(dm_build_1512, byte(len(dm_build_1513))) + dm_build_1511.Dm_build_1502(dm_build_1512+1, dm_build_1513)
}

func (dm_build_1515 *Dm_build_1279) Dm_build_1514(dm_build_1516 int, dm_build_1517 []byte) int {
	return dm_build_1515.Dm_build_1470(dm_build_1516, int16(len(dm_build_1517))) + dm_build_1515.Dm_build_1502(dm_build_1516+2, dm_build_1517)
}

func (dm_build_1519 *Dm_build_1279) Dm_build_1518(dm_build_1520 int, dm_build_1521 []byte) int {
	return dm_build_1519.Dm_build_1502(dm_build_1520, dm_build_1521) + dm_build_1519.Dm_build_1466(dm_build_1520+len(dm_build_1521), 0)
}

func (dm_build_1523 *Dm_build_1279) Dm_build_1522(dm_build_1524 int, dm_build_1525 string, dm_build_1526 string, dm_build_1527 *DmConnection) int {
	return dm_build_1523.Dm_build_1506(dm_build_1524, Dm_build_919.Dm_build_1135(dm_build_1525, dm_build_1526, dm_build_1527))
}

func (dm_build_1529 *Dm_build_1279) Dm_build_1528(dm_build_1530 int, dm_build_1531 string, dm_build_1532 string, dm_build_1533 *DmConnection) int {
	return dm_build_1529.Dm_build_1510(dm_build_1530, Dm_build_919.Dm_build_1135(dm_build_1531, dm_build_1532, dm_build_1533))
}

func (dm_build_1535 *Dm_build_1279) Dm_build_1534(dm_build_1536 int, dm_build_1537 string, dm_build_1538 string, dm_build_1539 *DmConnection) int {
	return dm_build_1535.Dm_build_1514(dm_build_1536, Dm_build_919.Dm_build_1135(dm_build_1537, dm_build_1538, dm_build_1539))
}

func (dm_build_1541 *Dm_build_1279) Dm_build_1540(dm_build_1542 int, dm_build_1543 string, dm_build_1544 string, dm_build_1545 *DmConnection) int {
	return dm_build_1541.Dm_build_1518(dm_build_1542, Dm_build_919.Dm_build_1135(dm_build_1543, dm_build_1544, dm_build_1545))
}

func (dm_build_1547 *Dm_build_1279) Dm_build_1546(dm_build_1548 int) byte {
	return Dm_build_919.Dm_build_1140(dm_build_1547.Dm_build_1573(dm_build_1548, 1))
}

func (dm_build_1550 *Dm_build_1279) Dm_build_1549(dm_build_1551 int) int16 {
	return Dm_build_919.Dm_build_1143(dm_build_1550.Dm_build_1573(dm_build_1551, 2))
}

func (dm_build_1553 *Dm_build_1279) Dm_build_1552(dm_build_1554 int) int32 {
	return Dm_build_919.Dm_build_1146(dm_build_1553.Dm_build_1573(dm_build_1554, 4))
}

func (dm_build_1556 *Dm_build_1279) Dm_build_1555(dm_build_1557 int) int64 {
	return Dm_build_919.Dm_build_1149(dm_build_1556.Dm_build_1573(dm_build_1557, 8))
}

func (dm_build_1559 *Dm_build_1279) Dm_build_1558(dm_build_1560 int) float32 {
	return Dm_build_919.Dm_build_1152(dm_build_1559.Dm_build_1573(dm_build_1560, 4))
}

func (dm_build_1562 *Dm_build_1279) Dm_build_1561(dm_build_1563 int) float64 {
	return Dm_build_919.Dm_build_1155(dm_build_1562.Dm_build_1573(dm_build_1563, 8))
}

func (dm_build_1565 *Dm_build_1279) Dm_build_1564(dm_build_1566 int) uint8 {
	return Dm_build_919.Dm_build_1158(dm_build_1565.Dm_build_1573(dm_build_1566, 1))
}

func (dm_build_1568 *Dm_build_1279) Dm_build_1567(dm_build_1569 int) uint16 {
	return Dm_build_919.Dm_build_1161(dm_build_1568.Dm_build_1573(dm_build_1569, 2))
}

func (dm_build_1571 *Dm_build_1279) Dm_build_1570(dm_build_1572 int) uint32 {
	return Dm_build_919.Dm_build_1164(dm_build_1571.Dm_build_1573(dm_build_1572, 4))
}

func (dm_build_1574 *Dm_build_1279) Dm_build_1573(dm_build_1575 int, dm_build_1576 int) []byte {
	return dm_build_1574.dm_build_1280[dm_build_1575 : dm_build_1575+dm_build_1576]
}

func (dm_build_1578 *Dm_build_1279) Dm_build_1577(dm_build_1579 int) []byte {
	dm_build_1580 := dm_build_1578.Dm_build_1552(dm_build_1579)
	return dm_build_1578.Dm_build_1573(dm_build_1579+4, int(dm_build_1580))
}

func (dm_build_1582 *Dm_build_1279) Dm_build_1581(dm_build_1583 int) []byte {
	dm_build_1584 := dm_build_1582.Dm_build_1546(dm_build_1583)
	return dm_build_1582.Dm_build_1573(dm_build_1583+1, int(dm_build_1584))
}

func (dm_build_1586 *Dm_build_1279) Dm_build_1585(dm_build_1587 int) []byte {
	dm_build_1588 := dm_build_1586.Dm_build_1549(dm_build_1587)
	return dm_build_1586.Dm_build_1573(dm_build_1587+2, int(dm_build_1588))
}

func (dm_build_1590 *Dm_build_1279) Dm_build_1589(dm_build_1591 int) []byte {
	dm_build_1592 := 0
	for dm_build_1590.Dm_build_1546(dm_build_1591) != 0 {
		dm_build_1591++
		dm_build_1592++
	}

	return dm_build_1590.Dm_build_1573(dm_build_1591-dm_build_1592, int(dm_build_1592))
}

func (dm_build_1594 *Dm_build_1279) Dm_build_1593(dm_build_1595 int, dm_build_1596 string, dm_build_1597 *DmConnection) string {
	return Dm_build_919.Dm_build_1172(dm_build_1594.Dm_build_1577(dm_build_1595), dm_build_1596, dm_build_1597)
}

func (dm_build_1599 *Dm_build_1279) Dm_build_1598(dm_build_1600 int, dm_build_1601 string, dm_build_1602 *DmConnection) string {
	return Dm_build_919.Dm_build_1172(dm_build_1599.Dm_build_1581(dm_build_1600), dm_build_1601, dm_build_1602)
}

func (dm_build_1604 *Dm_build_1279) Dm_build_1603(dm_build_1605 int, dm_build_1606 string, dm_build_1607 *DmConnection) string {
	return Dm_build_919.Dm_build_1172(dm_build_1604.Dm_build_1585(dm_build_1605), dm_build_1606, dm_build_1607)
}

func (dm_build_1609 *Dm_build_1279) Dm_build_1608(dm_build_1610 int, dm_build_1611 string, dm_build_1612 *DmConnection) string {
	return Dm_build_919.Dm_build_1172(dm_build_1609.Dm_build_1589(dm_build_1610), dm_build_1611, dm_build_1612)
}
