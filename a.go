/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/smilextay/dm/security"
)

const (
	Dm_build_0 = 8192
	Dm_build_1 = 2 * time.Second
)

type dm_build_2 struct {
	dm_build_3  *net.TCPConn
	dm_build_4  *tls.Conn
	dm_build_5  *Dm_build_1279
	dm_build_6  *DmConnection
	dm_build_7  security.Cipher
	dm_build_8  bool
	dm_build_9  bool
	dm_build_10 *security.DhKey

	dm_build_11 bool
	dm_build_12 string
	dm_build_13 bool
}

func dm_build_14(dm_build_15 *DmConnection) (*dm_build_2, error) {
	dm_build_16, dm_build_17 := dm_build_19(dm_build_15.dmConnector.host+":"+strconv.Itoa(int(dm_build_15.dmConnector.port)), time.Duration(dm_build_15.dmConnector.socketTimeout)*time.Second)
	if dm_build_17 != nil {
		return nil, dm_build_17
	}

	dm_build_18 := dm_build_2{}
	dm_build_18.dm_build_3 = dm_build_16
	dm_build_18.dm_build_5 = Dm_build_1282(Dm_build_284)
	dm_build_18.dm_build_6 = dm_build_15
	dm_build_18.dm_build_8 = false
	dm_build_18.dm_build_9 = false
	dm_build_18.dm_build_11 = false
	dm_build_18.dm_build_12 = ""
	dm_build_18.dm_build_13 = false
	dm_build_15.Access = &dm_build_18

	return &dm_build_18, nil
}

func dm_build_19(dm_build_20 string, dm_build_21 time.Duration) (*net.TCPConn, error) {
	dm_build_22, dm_build_23 := net.DialTimeout("tcp", dm_build_20, dm_build_21)
	if dm_build_23 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_20).throw()
	}

	if tcpConn, ok := dm_build_22.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_1)
		tcpConn.SetNoDelay(true)

		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_25 *dm_build_2) dm_build_24(dm_build_26 dm_build_405) bool {
	var dm_build_27 = dm_build_25.dm_build_6.dmConnector.compress
	if dm_build_26.dm_build_420() == Dm_build_312 || dm_build_27 == Dm_build_361 {
		return false
	}

	if dm_build_27 == Dm_build_359 {
		return true
	} else if dm_build_27 == Dm_build_360 {
		return !dm_build_25.dm_build_6.Local && dm_build_26.dm_build_418() > Dm_build_358
	}

	return false
}

func (dm_build_29 *dm_build_2) dm_build_28(dm_build_30 dm_build_405) bool {
	var dm_build_31 = dm_build_29.dm_build_6.dmConnector.compress
	if dm_build_30.dm_build_420() == Dm_build_312 || dm_build_31 == Dm_build_361 {
		return false
	}

	if dm_build_31 == Dm_build_359 {
		return true
	} else if dm_build_31 == Dm_build_360 {
		return dm_build_29.dm_build_5.Dm_build_1546(Dm_build_320) == 1
	}

	return false
}

func (dm_build_33 *dm_build_2) dm_build_32(dm_build_34 dm_build_405) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_36 := dm_build_34.dm_build_418()

	if dm_build_36 > 0 {

		if dm_build_33.dm_build_24(dm_build_34) {
			var retBytes, err = Compress(dm_build_33.dm_build_5, Dm_build_313, int(dm_build_36), int(dm_build_33.dm_build_6.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_33.dm_build_5.Dm_build_1293(Dm_build_313)

			dm_build_33.dm_build_5.Dm_build_1334(dm_build_36)

			dm_build_33.dm_build_5.Dm_build_1362(retBytes)

			dm_build_34.dm_build_419(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_33.dm_build_5.Dm_build_1466(Dm_build_320, 1)
		}

		if dm_build_33.dm_build_9 {
			dm_build_36 = dm_build_34.dm_build_418()
			var retBytes = dm_build_33.dm_build_7.Encrypt(dm_build_33.dm_build_5.Dm_build_1573(Dm_build_313, int(dm_build_36)), true)

			dm_build_33.dm_build_5.Dm_build_1293(Dm_build_313)

			dm_build_33.dm_build_5.Dm_build_1362(retBytes)

			dm_build_34.dm_build_419(int32(len(retBytes)))
		}
	}

	if dm_build_33.dm_build_5.Dm_build_1291() > Dm_build_285 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_34.dm_build_414()
	if dm_build_33.dm_build_267(dm_build_34) {
		if dm_build_33.dm_build_4 != nil {
			dm_build_33.dm_build_5.Dm_build_1296(0)
			if _, err := dm_build_33.dm_build_5.Dm_build_1315(dm_build_33.dm_build_4); err != nil {
				return err
			}
		}
	} else {
		dm_build_33.dm_build_5.Dm_build_1296(0)
		if _, err := dm_build_33.dm_build_5.Dm_build_1315(dm_build_33.dm_build_3); err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_38 *dm_build_2) dm_build_37(dm_build_39 dm_build_405) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_41 := int32(0)
	if dm_build_38.dm_build_267(dm_build_39) {
		if dm_build_38.dm_build_4 != nil {
			dm_build_38.dm_build_5.Dm_build_1293(0)
			if _, err := dm_build_38.dm_build_5.Dm_build_1309(dm_build_38.dm_build_4, Dm_build_313); err != nil {
				return err
			}

			dm_build_41 = dm_build_39.dm_build_418()
			if dm_build_41 > 0 {
				if _, err := dm_build_38.dm_build_5.Dm_build_1309(dm_build_38.dm_build_4, int(dm_build_41)); err != nil {
					return err
				}
			}
		}
	} else {

		dm_build_38.dm_build_5.Dm_build_1293(0)
		if _, err := dm_build_38.dm_build_5.Dm_build_1309(dm_build_38.dm_build_3, Dm_build_313); err != nil {
			return err
		}
		dm_build_41 = dm_build_39.dm_build_418()

		if dm_build_41 > 0 {
			if _, err := dm_build_38.dm_build_5.Dm_build_1309(dm_build_38.dm_build_3, int(dm_build_41)); err != nil {
				return err
			}
		}
	}

	dm_build_39.dm_build_415()

	dm_build_41 = dm_build_39.dm_build_418()
	if dm_build_41 <= 0 {
		return nil
	}

	if dm_build_38.dm_build_9 {
		ebytes := dm_build_38.dm_build_5.Dm_build_1573(Dm_build_313, int(dm_build_41))
		bytes, err := dm_build_38.dm_build_7.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_38.dm_build_5.Dm_build_1293(Dm_build_313)
		dm_build_38.dm_build_5.Dm_build_1362(bytes)
		dm_build_39.dm_build_419(int32(len(bytes)))
	}

	if dm_build_38.dm_build_28(dm_build_39) {

		dm_build_41 = dm_build_39.dm_build_418()
		cbytes := dm_build_38.dm_build_5.Dm_build_1573(Dm_build_313+ULINT_SIZE, int(dm_build_41-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_38.dm_build_6.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_38.dm_build_5.Dm_build_1293(Dm_build_313)
		dm_build_38.dm_build_5.Dm_build_1362(bytes)
		dm_build_39.dm_build_419(int32(len(bytes)))
	}
	return nil
}

func (dm_build_43 *dm_build_2) dm_build_42(dm_build_44 dm_build_405) (dm_build_45 interface{}, dm_build_46 error) {
	if dm_build_43.dm_build_13 {
		return nil, ECGO_CONNECTION_CLOSED.throw()
	}
	dm_build_47 := dm_build_43.dm_build_6
	dm_build_47.mu.Lock()
	defer dm_build_47.mu.Unlock()
	dm_build_46 = dm_build_44.dm_build_409(dm_build_44)
	if dm_build_46 != nil {
		return nil, dm_build_46
	}

	dm_build_46 = dm_build_43.dm_build_32(dm_build_44)
	if dm_build_46 != nil {
		return nil, dm_build_46
	}

	dm_build_46 = dm_build_43.dm_build_37(dm_build_44)
	if dm_build_46 != nil {
		return nil, dm_build_46
	}

	return dm_build_44.dm_build_413(dm_build_44)
}

func (dm_build_49 *dm_build_2) dm_build_48() (*dm_build_861, error) {

	Dm_build_50 := dm_build_867(dm_build_49)
	_, dm_build_51 := dm_build_49.dm_build_42(Dm_build_50)
	if dm_build_51 != nil {
		return nil, dm_build_51
	}

	return Dm_build_50, nil
}

func (dm_build_53 *dm_build_2) dm_build_52() error {

	dm_build_54 := dm_build_728(dm_build_53)
	_, dm_build_55 := dm_build_53.dm_build_42(dm_build_54)
	if dm_build_55 != nil {
		return dm_build_55
	}

	return nil
}

func (dm_build_57 *dm_build_2) dm_build_56() error {

	var dm_build_58 *dm_build_861
	var err error
	if dm_build_58, err = dm_build_57.dm_build_48(); err != nil {
		return err
	}

	if dm_build_57.dm_build_6.sslEncrypt == 2 {
		if err = dm_build_57.dm_build_263(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_57.dm_build_6.sslEncrypt == 1 {
		if err = dm_build_57.dm_build_263(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_57.dm_build_9 || dm_build_57.dm_build_8 {
		k, err := dm_build_57.dm_build_253()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_58.Dm_build_865)
		encryptType := dm_build_58.dm_build_863
		hashType := int(dm_build_58.Dm_build_864)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_57.dm_build_256(encryptType, sessionKey, dm_build_57.dm_build_6.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_57.dm_build_52(); err != nil {
		return err
	}
	return nil
}

func (dm_build_61 *dm_build_2) Dm_build_60(dm_build_62 *DmStatement) error {
	dm_build_63 := dm_build_890(dm_build_61, dm_build_62)
	_, dm_build_64 := dm_build_61.dm_build_42(dm_build_63)
	if dm_build_64 != nil {
		return dm_build_64
	}

	return nil
}

func (dm_build_66 *dm_build_2) Dm_build_65(dm_build_67 int32) error {
	dm_build_68 := dm_build_900(dm_build_66, dm_build_67)
	_, dm_build_69 := dm_build_66.dm_build_42(dm_build_68)
	if dm_build_69 != nil {
		return dm_build_69
	}

	return nil
}

func (dm_build_71 *dm_build_2) Dm_build_70(dm_build_72 *DmStatement, dm_build_73 bool, dm_build_74 int16) (*execRetInfo, error) {
	dm_build_75 := dm_build_767(dm_build_71, dm_build_72, dm_build_73, dm_build_74)
	dm_build_76, dm_build_77 := dm_build_71.dm_build_42(dm_build_75)
	if dm_build_77 != nil {
		return nil, dm_build_77
	}
	return dm_build_76.(*execRetInfo), nil
}

func (dm_build_79 *dm_build_2) Dm_build_78(dm_build_80 *DmStatement, dm_build_81 int16) (*execRetInfo, error) {
	return dm_build_79.Dm_build_70(dm_build_80, false, Dm_build_365)
}

func (dm_build_83 *dm_build_2) Dm_build_82(dm_build_84 *DmStatement, dm_build_85 []OptParameter) (*execRetInfo, error) {
	dm_build_86, dm_build_87 := dm_build_83.dm_build_42(dm_build_508(dm_build_83, dm_build_84, dm_build_85))
	if dm_build_87 != nil {
		return nil, dm_build_87
	}

	return dm_build_86.(*execRetInfo), nil
}

func (dm_build_89 *dm_build_2) Dm_build_88(dm_build_90 *DmStatement, dm_build_91 int16) (*execRetInfo, error) {
	return dm_build_89.Dm_build_70(dm_build_90, true, dm_build_91)
}

func (dm_build_93 *dm_build_2) Dm_build_92(dm_build_94 *DmStatement, dm_build_95 [][]interface{}) (*execRetInfo, error) {
	dm_build_96 := dm_build_540(dm_build_93, dm_build_94, dm_build_95)
	dm_build_97, dm_build_98 := dm_build_93.dm_build_42(dm_build_96)
	if dm_build_98 != nil {
		return nil, dm_build_98
	}
	return dm_build_97.(*execRetInfo), nil
}

func (dm_build_100 *dm_build_2) Dm_build_99(dm_build_101 *DmStatement, dm_build_102 [][]interface{}, dm_build_103 bool) (*execRetInfo, error) {
	var dm_build_104, dm_build_105 = 0, 0
	var dm_build_106 = len(dm_build_102)
	var dm_build_107 [][]interface{}
	var dm_build_108 = NewExceInfo()
	dm_build_108.updateCounts = make([]int64, dm_build_106)
	var dm_build_109 = false
	for dm_build_104 < dm_build_106 {
		for dm_build_105 = dm_build_104; dm_build_105 < dm_build_106; dm_build_105++ {
			paramData := dm_build_102[dm_build_105]
			bindData := make([]interface{}, dm_build_101.paramCount)
			dm_build_109 = false
			for icol := 0; icol < int(dm_build_101.paramCount); icol++ {
				if dm_build_101.bindParams[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_100.dm_build_236(bindData, paramData, icol) {
					dm_build_109 = true
					break
				}
			}

			if dm_build_109 {
				break
			}
			dm_build_107 = append(dm_build_107, bindData)
		}

		if dm_build_105 != dm_build_104 {
			tmpExecInfo, err := dm_build_100.Dm_build_92(dm_build_101, dm_build_107)
			if err != nil {
				return nil, err
			}
			dm_build_107 = dm_build_107[0:0]
			dm_build_108.union(tmpExecInfo, dm_build_104, dm_build_105-dm_build_104)
		}

		if dm_build_105 < dm_build_106 {
			tmpExecInfo, err := dm_build_100.Dm_build_110(dm_build_101, dm_build_102[dm_build_105], dm_build_103)
			if err != nil {
				return nil, err
			}

			dm_build_103 = true
			dm_build_108.union(tmpExecInfo, dm_build_105, 1)
		}

		dm_build_104 = dm_build_105 + 1
	}
	for _, i := range dm_build_108.updateCounts {
		if i > 0 {
			dm_build_108.updateCount += i
		}
	}
	return dm_build_108, nil
}

func (dm_build_111 *dm_build_2) Dm_build_110(dm_build_112 *DmStatement, dm_build_113 []interface{}, dm_build_114 bool) (*execRetInfo, error) {

	var dm_build_115 = make([]interface{}, dm_build_112.paramCount)
	for icol := 0; icol < int(dm_build_112.paramCount); icol++ {
		if dm_build_112.bindParams[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_111.dm_build_236(dm_build_115, dm_build_113, icol) {

			if !dm_build_114 {
				preExecute := dm_build_756(dm_build_111, dm_build_112, dm_build_112.bindParams)
				dm_build_111.dm_build_42(preExecute)
				dm_build_114 = true
			}

			dm_build_111.dm_build_242(dm_build_112, dm_build_112.bindParams[icol], icol, dm_build_113[icol].(iOffRowBinder))
			dm_build_115[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_116 = make([][]interface{}, 1, 1)
	dm_build_116[0] = dm_build_115

	dm_build_117 := dm_build_540(dm_build_111, dm_build_112, dm_build_116)
	dm_build_118, dm_build_119 := dm_build_111.dm_build_42(dm_build_117)
	if dm_build_119 != nil {
		return nil, dm_build_119
	}
	return dm_build_118.(*execRetInfo), nil
}

func (dm_build_121 *dm_build_2) Dm_build_120(dm_build_122 *DmStatement, dm_build_123 int16) (*execRetInfo, error) {
	dm_build_124 := dm_build_743(dm_build_121, dm_build_122, dm_build_123)

	dm_build_125, dm_build_126 := dm_build_121.dm_build_42(dm_build_124)
	if dm_build_126 != nil {
		return nil, dm_build_126
	}
	return dm_build_125.(*execRetInfo), nil
}

func (dm_build_128 *dm_build_2) Dm_build_127(dm_build_129 *innerRows, dm_build_130 int64) (*execRetInfo, error) {
	dm_build_131 := dm_build_646(dm_build_128, dm_build_129, dm_build_130, INT64_MAX)
	dm_build_132, dm_build_133 := dm_build_128.dm_build_42(dm_build_131)
	if dm_build_133 != nil {
		return nil, dm_build_133
	}
	return dm_build_132.(*execRetInfo), nil
}

func (dm_build_135 *dm_build_2) Commit() error {
	dm_build_136 := dm_build_493(dm_build_135)
	_, dm_build_137 := dm_build_135.dm_build_42(dm_build_136)
	if dm_build_137 != nil {
		return dm_build_137
	}

	return nil
}

func (dm_build_139 *dm_build_2) Rollback() error {
	dm_build_140 := dm_build_805(dm_build_139)
	_, dm_build_141 := dm_build_139.dm_build_42(dm_build_140)
	if dm_build_141 != nil {
		return dm_build_141
	}

	return nil
}

func (dm_build_143 *dm_build_2) Dm_build_142(dm_build_144 *DmConnection) error {
	dm_build_145 := dm_build_810(dm_build_143, dm_build_144.IsoLevel)
	_, dm_build_146 := dm_build_143.dm_build_42(dm_build_145)
	if dm_build_146 != nil {
		return dm_build_146
	}

	return nil
}

func (dm_build_148 *dm_build_2) Dm_build_147(dm_build_149 *DmStatement, dm_build_150 string) error {
	dm_build_151 := dm_build_498(dm_build_148, dm_build_149, dm_build_150)
	_, dm_build_152 := dm_build_148.dm_build_42(dm_build_151)
	if dm_build_152 != nil {
		return dm_build_152
	}

	return nil
}

func (dm_build_154 *dm_build_2) Dm_build_153(dm_build_155 []uint32) ([]int64, error) {
	dm_build_156 := dm_build_908(dm_build_154, dm_build_155)
	dm_build_157, dm_build_158 := dm_build_154.dm_build_42(dm_build_156)
	if dm_build_158 != nil {
		return nil, dm_build_158
	}
	return dm_build_157.([]int64), nil
}

func (dm_build_160 *dm_build_2) Close() error {
	if dm_build_160.dm_build_13 {
		return nil
	}

	dm_build_161 := dm_build_160.dm_build_3.Close()
	if dm_build_161 != nil {
		return dm_build_161
	}

	dm_build_160.dm_build_6 = nil
	dm_build_160.dm_build_13 = true
	return nil
}

func (dm_build_163 *dm_build_2) dm_build_162(dm_build_164 *lob) (int64, error) {
	dm_build_165 := dm_build_679(dm_build_163, dm_build_164)
	dm_build_166, dm_build_167 := dm_build_163.dm_build_42(dm_build_165)
	if dm_build_167 != nil {
		return 0, dm_build_167
	}
	return dm_build_166.(int64), nil
}

func (dm_build_169 *dm_build_2) dm_build_168(dm_build_170 *lob, dm_build_171 int32, dm_build_172 int32) (*lobRetInfo, error) {
	dm_build_173 := dm_build_664(dm_build_169, dm_build_170, int(dm_build_171), int(dm_build_172))
	dm_build_174, dm_build_175 := dm_build_169.dm_build_42(dm_build_173)
	if dm_build_175 != nil {
		return nil, dm_build_175
	}
	return dm_build_174.(*lobRetInfo), nil
}

func (dm_build_177 *dm_build_2) dm_build_176(dm_build_178 *DmBlob, dm_build_179 int32, dm_build_180 int32) ([]byte, error) {
	var dm_build_181 = make([]byte, dm_build_180)
	var dm_build_182 int32 = 0
	var dm_build_183 int32 = 0
	var dm_build_184 *lobRetInfo
	var dm_build_185 []byte
	var dm_build_186 error
	for dm_build_182 < dm_build_180 {
		dm_build_183 = dm_build_180 - dm_build_182
		if dm_build_183 > Dm_build_398 {
			dm_build_183 = Dm_build_398
		}
		dm_build_184, dm_build_186 = dm_build_177.dm_build_168(&dm_build_178.lob, dm_build_179+dm_build_182, dm_build_183)
		if dm_build_186 != nil {
			return nil, dm_build_186
		}
		dm_build_185 = dm_build_184.data
		if dm_build_185 == nil || len(dm_build_185) == 0 {
			break
		}
		Dm_build_919.Dm_build_975(dm_build_181, int(dm_build_182), dm_build_185, 0, len(dm_build_185))
		dm_build_182 += int32(len(dm_build_185))
		if dm_build_178.readOver {
			break
		}
	}
	return dm_build_181, nil
}

func (dm_build_188 *dm_build_2) dm_build_187(dm_build_189 *DmClob, dm_build_190 int32, dm_build_191 int32) (string, error) {
	var dm_build_192 bytes.Buffer
	var dm_build_193 int32 = 0
	var dm_build_194 int32 = 0
	var dm_build_195 *lobRetInfo
	var dm_build_196 []byte
	var dm_build_197 string
	var dm_build_198 error
	for dm_build_193 < dm_build_191 {
		dm_build_194 = dm_build_191 - dm_build_193
		if dm_build_194 > Dm_build_398/2 {
			dm_build_194 = Dm_build_398 / 2
		}
		dm_build_195, dm_build_198 = dm_build_188.dm_build_168(&dm_build_189.lob, dm_build_190+dm_build_193, dm_build_194)
		if dm_build_198 != nil {
			return "", dm_build_198
		}
		dm_build_196 = dm_build_195.data
		if dm_build_196 == nil || len(dm_build_196) == 0 {
			break
		}
		dm_build_197 = Dm_build_919.Dm_build_1076(dm_build_196, 0, len(dm_build_196), dm_build_189.serverEncoding, dm_build_188.dm_build_6)

		dm_build_192.WriteString(dm_build_197)
		var strLen = dm_build_195.charLen
		if strLen == -1 {
			strLen = int64(utf8.RuneCountInString(dm_build_197))
		}
		dm_build_193 += int32(strLen)
		if dm_build_189.readOver {
			break
		}
	}
	return dm_build_192.String(), nil
}

func (dm_build_200 *dm_build_2) dm_build_199(dm_build_201 *DmClob, dm_build_202 int, dm_build_203 string, dm_build_204 string) (int, error) {
	var dm_build_205 = Dm_build_919.Dm_build_1135(dm_build_203, dm_build_204, dm_build_200.dm_build_6)
	var dm_build_206 = 0
	var dm_build_207 = len(dm_build_205)
	var dm_build_208 = 0
	var dm_build_209 = 0
	var dm_build_210 = 0
	var dm_build_211 = dm_build_207/Dm_build_397 + 1
	var dm_build_212 byte = 0
	var dm_build_213 byte = 0x01
	var dm_build_214 byte = 0x02
	for i := 0; i < dm_build_211; i++ {
		dm_build_212 = 0
		if i == 0 {
			dm_build_212 |= dm_build_213
		}
		if i == dm_build_211-1 {
			dm_build_212 |= dm_build_214
		}
		dm_build_210 = dm_build_207 - dm_build_209
		if dm_build_210 > Dm_build_397 {
			dm_build_210 = Dm_build_397
		}

		setLobData := dm_build_824(dm_build_200, &dm_build_201.lob, dm_build_212, dm_build_202, dm_build_205, dm_build_206, dm_build_210)
		ret, err := dm_build_200.dm_build_42(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_208, nil
		} else {
			dm_build_202 += int(tmp)
			dm_build_208 += int(tmp)
			dm_build_209 += dm_build_210
			dm_build_206 += dm_build_210
		}
	}
	return dm_build_208, nil
}

func (dm_build_216 *dm_build_2) dm_build_215(dm_build_217 *DmBlob, dm_build_218 int, dm_build_219 []byte) (int, error) {
	var dm_build_220 = 0
	var dm_build_221 = len(dm_build_219)
	var dm_build_222 = 0
	var dm_build_223 = 0
	var dm_build_224 = 0
	var dm_build_225 = dm_build_221/Dm_build_397 + 1
	var dm_build_226 byte = 0
	var dm_build_227 byte = 0x01
	var dm_build_228 byte = 0x02
	for i := 0; i < dm_build_225; i++ {
		dm_build_226 = 0
		if i == 0 {
			dm_build_226 |= dm_build_227
		}
		if i == dm_build_225-1 {
			dm_build_226 |= dm_build_228
		}
		dm_build_224 = dm_build_221 - dm_build_223
		if dm_build_224 > Dm_build_397 {
			dm_build_224 = Dm_build_397
		}

		setLobData := dm_build_824(dm_build_216, &dm_build_217.lob, dm_build_226, dm_build_218, dm_build_219, dm_build_220, dm_build_224)
		ret, err := dm_build_216.dm_build_42(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_222, nil
		} else {
			dm_build_218 += int(tmp)
			dm_build_222 += int(tmp)
			dm_build_223 += dm_build_224
			dm_build_220 += dm_build_224
		}
	}
	return dm_build_222, nil
}

func (dm_build_230 *dm_build_2) dm_build_229(dm_build_231 *lob, dm_build_232 int) (int64, error) {
	dm_build_233 := dm_build_690(dm_build_230, dm_build_231, dm_build_232)
	dm_build_234, dm_build_235 := dm_build_230.dm_build_42(dm_build_233)
	if dm_build_235 != nil {
		return dm_build_231.length, dm_build_235
	}
	return dm_build_234.(int64), nil
}

func (dm_build_237 *dm_build_2) dm_build_236(dm_build_238 []interface{}, dm_build_239 []interface{}, dm_build_240 int) bool {
	var dm_build_241 = false
	dm_build_238[dm_build_240] = dm_build_239[dm_build_240]

	if binder, ok := dm_build_239[dm_build_240].(iOffRowBinder); ok {
		dm_build_241 = true
		dm_build_238[dm_build_240] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_237.dm_build_6) {
			dm_build_238[dm_build_240] = &lobCtl{lob.buildCtlData()}
			dm_build_241 = false
		}
	} else {
		dm_build_238[dm_build_240] = dm_build_239[dm_build_240]
	}
	return dm_build_241
}

func (dm_build_243 *dm_build_2) dm_build_242(dm_build_244 *DmStatement, dm_build_245 parameter, dm_build_246 int, dm_build_247 iOffRowBinder) error {
	var dm_build_248 = Dm_build_1205()
	dm_build_247.read(dm_build_248)
	var dm_build_249 = 0
	for !dm_build_247.isReadOver() || dm_build_248.Dm_build_1206() > 0 {
		if !dm_build_247.isReadOver() && dm_build_248.Dm_build_1206() < Dm_build_397 {
			dm_build_247.read(dm_build_248)
		}
		if dm_build_248.Dm_build_1206() > Dm_build_397 {
			dm_build_249 = Dm_build_397
		} else {
			dm_build_249 = dm_build_248.Dm_build_1206()
		}

		putData := dm_build_795(dm_build_243, dm_build_244, int16(dm_build_246), dm_build_248, int32(dm_build_249))
		_, err := dm_build_243.dm_build_42(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_251 *dm_build_2) dm_build_250() ([]byte, error) {
	var dm_build_252 error
	if dm_build_251.dm_build_10 == nil {
		if dm_build_251.dm_build_10, dm_build_252 = security.NewClientKeyPair(); dm_build_252 != nil {
			return nil, dm_build_252
		}
	}
	return security.Bn2Bytes(dm_build_251.dm_build_10.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_254 *dm_build_2) dm_build_253() (*security.DhKey, error) {
	var dm_build_255 error
	if dm_build_254.dm_build_10 == nil {
		if dm_build_254.dm_build_10, dm_build_255 = security.NewClientKeyPair(); dm_build_255 != nil {
			return nil, dm_build_255
		}
	}
	return dm_build_254.dm_build_10, nil
}

func (dm_build_257 *dm_build_2) dm_build_256(dm_build_258 int, dm_build_259 []byte, dm_build_260 string, dm_build_261 int) (dm_build_262 error) {
	if dm_build_258 > 0 && dm_build_258 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_259 != nil {
		dm_build_257.dm_build_7, dm_build_262 = security.NewSymmCipher(dm_build_258, dm_build_259)
	} else if dm_build_258 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_257.dm_build_7, dm_build_262 = security.NewThirdPartCipher(dm_build_258, dm_build_259, dm_build_260, dm_build_261); dm_build_262 != nil {
			dm_build_262 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_262.Error()).throw()
		}
	}
	return
}

func (dm_build_264 *dm_build_2) dm_build_263(dm_build_265 bool) (dm_build_266 error) {
	if dm_build_264.dm_build_4, dm_build_266 = security.NewTLSFromTCP(dm_build_264.dm_build_3, dm_build_264.dm_build_6.dmConnector.sslCertPath, dm_build_264.dm_build_6.dmConnector.sslKeyPath, dm_build_264.dm_build_6.dmConnector.user); dm_build_266 != nil {
		return
	}
	if !dm_build_265 {
		dm_build_264.dm_build_4 = nil
	}
	return
}

func (dm_build_268 *dm_build_2) dm_build_267(dm_build_269 dm_build_405) bool {
	return dm_build_269.dm_build_420() != Dm_build_312 && dm_build_268.dm_build_6.sslEncrypt == 1
}
