// Code generated by "esc -o store/packed_migrations.go -pkg store -modtime 1 store/migrations"; DO NOT EDIT.

package store

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/store/migrations/1533636853_initial.down.sql": {
		name:    "1533636853_initial.down.sql",
		local:   "store/migrations/1533636853_initial.down.sql",
		size:    87,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicgnyD1AIcXTycVUoKC3OiE8tS80rQRVPzs/NxRAsSi3LTC2HK3f29/X1
DLHmAgQAAP//Z/tze1cAAAA=
`,
	},

	"/store/migrations/1533636853_initial.up.sql": {
		name:    "1533636853_initial.up.sql",
		local:   "store/migrations/1533636853_initial.up.sql",
		size:    945,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/8ySTW7CMBCF1/EpZgkSN2AFyK1QIVRpumAVOfEQpkrsyB7Tn9NXQWoVcCraXbfvfVH0
/M1S3q/TuRCrTC5yCfliuZHg8ET4WuAJDcNEJKQhBNKQ7nJInzcbeMzW20W2hwe5n4nEs+LggfGNv5GZ
SDpnT6TRRQUZRmdUU5COusqhYtSFYmBq0bNqO/4YEqHTNwjyRYuuRlU2CKW1DSoz7L0NrkJ48daUw/z8
URxX1hyoDk4xWRPXDjvria177/eUVJO5WGRCW6IbKUrlR/52RKWvUjGdiytDlW3b38kZuiy+2EzeyUym
K/l04XpCejoTyYEajMQ0ZBB6czW6YX7mIov9k2k0FV7vHtnSBX/8p7fWPzOxH5GnyTOZioufkRt381f9
u+12nc/FZwAAAP//zPqiPbEDAAA=
`,
	},

	"/store/migrations/1537263282_add_review_target.down.sql": {
		name:    "1537263282_add_review_target.down.sql",
		local:   "store/migrations/1537263282_add_review_target.down.sql",
		size:    99,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicgnyD1AIcXTycVUoSi3LTC2PL0ksSk8tsebicvQJcQ1ClUstS80rUQDr
cfb3CfX1Q9UUn5lizcXl7O/r6xlizQUIAAD//5hwoGFjAAAA
`,
	},

	"/store/migrations/1537263282_add_review_target.up.sql": {
		name:    "1537263282_add_review_target.up.sql",
		local:   "store/migrations/1537263282_add_review_target.up.sql",
		size:    279,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/2TPwUoDMRDG8fPOU8yxQt9gT2k6ymKSlZgeelpaMpQBzZYxu+rbSw8KseffxzD/HT0N
oQewkUwiTGbnCJVX4c+pnvTCFTfQScZlkYxhTBgOzuFLHLyJR3ym4xa6q86rZFas/FX/RlvopFTWcnqb
JN+Z8nX+kDrr903PcpHSeFnez6z/AR56ADAuUWyf5ZVLRbPfox3dwYc2YvotiPRIkYKl13awkXy7bEfv
h9TDTwAAAP//hNuj3xcBAAA=
`,
	},

	"/store/migrations/1537263364_add_analyzer_name_to_comment.down.sql": {
		name:    "1537263364_add_analyzer_name_to_comment.down.sql",
		local:   "store/migrations/1537263364_add_analyzer_name_to_comment.down.sql",
		size:    59,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicvQJcQ1SCHF08nFVSM7PzU3NK1FwCfIPUHD29wn19VNIzEvMqaxKLbLm
4nL29/X1DLHmAgQAAP//djugFzsAAAA=
`,
	},

	"/store/migrations/1537263364_add_analyzer_name_to_comment.up.sql": {
		name:    "1537263364_add_analyzer_name_to_comment.up.sql",
		local:   "store/migrations/1537263364_add_analyzer_name_to_comment.up.sql",
		size:    83,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/wTAQQoCMQwF0H1O8XdziK46M0UKaQqSHqBoXLUVJIJ6et+eLlkCUWRNV2jcOeH2nNOW
I54njsqtCPrq4/uzF9w+DqkKacy426O/h2PbAtFRS8ka6B8AAP//Ux0Zt1MAAAA=
`,
	},

	"/store/migrations/1537266831_add_review_old_id.down.sql": {
		name:    "1537266831_add_review_old_id.down.sql",
		local:   "store/migrations/1537266831_add_review_old_id.down.sql",
		size:    71,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicvQJcQ1SCHF08nFVKEoty0wtj08tS80rUXAJ8g9QcPb3CfX1U8jPSYnP
zCtJLcpLzInPTLHm4nL29/X1DLHmAgQAAP//0vNoAUcAAAA=
`,
	},

	"/store/migrations/1537266831_add_review_old_id.up.sql": {
		name:    "1537266831_add_review_old_id.up.sql",
		local:   "store/migrations/1537266831_add_review_old_id.up.sql",
		size:    95,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/wTAwQrDIAwG4Hue4r/1ITzZVkYhKoz0LAUzEMRBydwef98aHkdyRJ4lPCF+5YBbZ9Nv
0anD4PcdW+YzJrx7LW2Y3uPqpVWY/gwpC9LJjKqv69MNy+KIthzjIY7+AQAA//9qFuu3XwAAAA==
`,
	},

	"/store/migrations/1537267179_add_timestamps_to_comments_and_target.down.sql": {
		name:    "1537267179_add_timestamps_to_comments_and_target.down.sql",
		local:   "store/migrations/1537267179_add_timestamps_to_comments_and_target.down.sql",
		size:    208,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicvQJcQ1SCHF08nFVSM7PzU3NK1FwCfIPUHD29wn19VNILkpNLElNiU8s
IUJxaUEKdsVFqWWZqeXxJYlF6alEmo9bC4otzv6+vp4h1lyAAAAA///vMdEl0AAAAA==
`,
	},

	"/store/migrations/1537267179_add_timestamps_to_comments_and_target.up.sql": {
		name:    "1537267179_add_timestamps_to_comments_and_target.up.sql",
		local:   "store/migrations/1537267179_add_timestamps_to_comments_and_target.up.sql",
		size:    490,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/7TPsWrDMBSF4V1PcbZA2hB5Kq0nJ3ZLQLahyHO42NetQYqCdF3TPn0he+niwJnPx3+o
3k5NrtR+izkxBh5pdgJKSEJREEbIJ0Mmz4+YBMvkHGgcuReEi/tGcANiWBK2e6UKY6t32OJgKvTBe74I
irLEsTVd3aCPTMLDmeR2mIT8VX7QtBZNZwzK6rXojMUme37SO53tdAatX2570HqT/yvM12FNIfLXxMtZ
KH7wPUv+dFboObZ1fbK5+g0AAP//WzsWhuoBAAA=
`,
	},

	"/store/migrations/1537268276_data_migrate.down.sql": {
		name:    "1537268276_data_migrate.down.sql",
		local:   "store/migrations/1537268276_data_migrate.down.sql",
		size:    73,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/yTLsQ3DIBAF0D5T/A6Jhpk+4QSnKHcSfBfe3oUHeK1iJCK1PCa6fXkdg6scDIr4+9yU
Z0CLwkg7UYS+jT8w7tfV9nkCAAD//7QpIAJJAAAA
`,
	},

	"/store/migrations/1537268276_data_migrate.up.sql": {
		name:    "1537268276_data_migrate.up.sql",
		local:   "store/migrations/1537268276_data_migrate.up.sql",
		size:    1324,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/6xTwW7aQBC9+yueuJAgGnroKVEOBJbWkrEre1HTk7WwQ2pp2UXrNZS/r9aGyI5ppai5
2TNv3szbN/PEvobxQxBMRkjpUNCRHUi7e2wLpZAoGWpHVgsVzjGaBKvv8ylnsDUyJw9FxjiMknlxRuaF
xCPafz++sZRdwQyHTePICImqKuQnU5Z70G9HuiyMhjNYE8Rakf98IU1WOKqhfppZyvw07JmzOAuTGOEC
ccLBnsOMZxi8Ug7aArmwL+TusbHkybbG+tHQiBlNgjDOWMoRxjy5CHV1CW4KOcZgb82hkGQH47bIMSzt
TVk4Y0/170BXu3WNahrJXLgxqr08f98GGYvYjNdq8ou2/PDl5vbjmmCRJsuOXUFjRkeYdyPMoCulrmxC
Sa4Pv74Lgd+FHvaxG7orZNAeq4k2cwUAOox3l4fosbwmpvG8X9ZdtDf9W7mrxZ1H7pV3s1cJGld6lefw
X3r+yxAhJSRtRaUcDkJVVPqD8Gu7Mara6XKMNW1EVRKOBGn00OFoC0dwv2gHoU87YwlCSx84YSM8Yk11
C2/mNOIsBZ8+RQxJHP3snniTnSXRahm3drM+/TlbTFcRr4/5PSwtF/6PqONHl+rz+5jOx/SWIpgly2XI
H4I/AQAA//8OZoMGLAUAAA==
`,
	},

	"/store/migrations/1537455097_delete_old_columns.down.sql": {
		name:    "1537455097_delete_old_columns.down.sql",
		local:   "store/migrations/1537455097_delete_old_columns.down.sql",
		size:    270,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/5zMQarCMBAA0H1OMffoKm3DpzBJ4JOui6WDDOikjGPU24t7F+oB3uvD35Q65zyW8A/F
9xhAqTHdFmokBn4cYcg4xwS71sYbKRjdDVIukGbEDzGLkcrhtPD2k1fa64Wt6uM1rHxk+fqQ63klfYOH
HONUOvcMAAD//7TEya4OAQAA
`,
	},

	"/store/migrations/1537455097_delete_old_columns.up.sql": {
		name:    "1537455097_delete_old_columns.up.sql",
		local:   "store/migrations/1537455097_delete_old_columns.up.sql",
		size:    214,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicvQJcQ1SCHF08nFVKEoty0wtj08tS80rUXAJ8g9QcPb3CfX1Uygoyi/L
TEktIlZ9Zl5JalFeYk58ZgqxWopSC/KLM0vyiypJ0JRXmpsEdpWzv6+vZ4g1FyAAAP//t7XXZtYAAAA=
`,
	},

	"/store/migrations/1537528734_review_internal_id.down.sql": {
		name:    "1537528734_review_internal_id.down.sql",
		local:   "store/migrations/1537528734_review_internal_id.down.sql",
		size:    67,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicvQJcQ1SCHF08nFVKEoty0wtj08tS80rUXAJ8g9QcPb3CfX1U8jMK0kt
ykvMic9Msebicvb39fUMseYCBAAA//9aEbrDQwAAAA==
`,
	},

	"/store/migrations/1537528734_review_internal_id.up.sql": {
		name:    "1537528734_review_internal_id.up.sql",
		local:   "store/migrations/1537528734_review_internal_id.up.sql",
		size:    91,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/wTAwQqDMAwG4Hue4r/5ED1VLUNIUxjxLIIZFKSHknV7fL85vTYJRJE1vaFx5oRuo9rv
sGHNEdcVS+E9C2pz6+28j3rB7e+QopCdGZd9zu/tmKZAtJScNw30BAAA//+NTHcJWwAAAA==
`,
	},

	"/store/migrations/1537528735_migrate_review_event_id.down.sql": {
		name:    "1537528735_migrate_review_event_id.down.sql",
		local:   "store/migrations/1537528735_migrate_review_event_id.down.sql",
		size:    0,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/wEAAP//AAAAAAAAAAA=
`,
	},

	"/store/migrations/1537528735_migrate_review_event_id.up.sql": {
		name:    "1537528735_migrate_review_event_id.up.sql",
		local:   "store/migrations/1537528735_migrate_review_event_id.up.sql",
		size:    278,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/1SQwUrEMBCG73mKYS/pQi143dKF6mY1h6bSRNxbCc3QRCQpaagKPrx0UejehmG++T/+
B/bERUnIY8dqxYBdFBOStwL4GUSrgF24VBJ20zjE7ymFXUnI68tpvY24OPzscUGfoJaARDIFzieMXn/0
zlToh2AwM27EOWVD8INOWSqmGBZnMOZAf2gOqdgwfzssLGpzOLzPwd8dj9Tq2dJ9DnS2+v46WPyie3Lu
2uZfJOk44tUkkbdn1jHA7eeKUqjFaY0zUAEWN1jvzNpC2zRcleQ3AAD//2L4P4IWAQAA
`,
	},

	"/store/migrations/1539095669_delete_old_internal_id.down.sql": {
		name:    "1539095669_delete_old_internal_id.down.sql",
		local:   "store/migrations/1539095669_delete_old_internal_id.down.sql",
		size:    84,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/wTAwQrDIAwG4Hue4n8PT7aVUYgRRnqWQXMQxIEEt8fvt6XXKYEosqY3NG6cMG01+1Vb
NhzxOLAXvrLg2+/ahtscn17bDbe/Q4pCLuZAtJecTw30BAAA//94UwlHVAAAAA==
`,
	},

	"/store/migrations/1539095669_delete_old_internal_id.up.sql": {
		name:    "1539095669_delete_old_internal_id.up.sql",
		local:   "store/migrations/1539095669_delete_old_internal_id.up.sql",
		size:    71,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicvQJcQ1SCHF08nFVKEoty0wtj08tS80rUXAJ8g9QcPb3CfX1U8jPSYnP
zCtJLcpLzInPTLHm4nL29/X1DLHmAgQAAP//0vNoAUcAAAA=
`,
	},

	"/store/migrations/1547744983_organizations.down.sql": {
		name:    "1547744983_organizations.down.sql",
		local:   "store/migrations/1547744983_organizations.down.sql",
		size:    42,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicgnyD1AIcXTycVXIL0pPzMusSizJzM+z5uJy9vf19Qyx5gIEAAD//3ZC
3tcqAAAA
`,
	},

	"/store/migrations/1547744983_organizations.up.sql": {
		name:    "1547744983_organizations.up.sql",
		local:   "store/migrations/1547744983_organizations.up.sql",
		size:    155,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/2TKwQoCIRCA4bPzFHMs2Dfw5C4SkrohdthTSNoyEGOIG9HTdwui6/9/oz4YLwGmoFXU
GNVoNda2JqZ36lQZdyAo47ZRRj9H9Gdr8RSMU2HBo14GEI9Wn5RLw15e/YsGEMS9NE73C+W/d618o/U3
w14CwDQ7Z6KETwAAAP//+lQfNZsAAAA=
`,
	},

	"/store/migrations/1548157581_organizations_constraints.down.sql": {
		name:    "1548157581_organizations_constraints.down.sql",
		local:   "store/migrations/1548157581_organizations_constraints.down.sql",
		size:    40,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3IJ8g9Q8PRzcY1QyC9KT8zLrEosyczPi0/Ozy3IL84sSY0vyE6ttOYCBAAA//+4y7zd
KAAAAA==
`,
	},

	"/store/migrations/1548157581_organizations_constraints.up.sql": {
		name:    "1548157581_organizations_constraints.up.sql",
		local:   "store/migrations/1548157581_organizations_constraints.up.sql",
		size:    218,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3SNsW7CMBRF5/or7thGqfIBnarWQ5ZURURiix72IzzFsS3bQYSvR8oEA/O955ymwhg+
J3KOrjCJqbAFQSyWRSxikpnSionXGseloJwZFKMTQ0WChwujGJApWVGGnLbDAwXJeI8pXMRyqiG+cPLk
BrEfqBqlfnb6e6/Rd+1/r9F2v/qAkEbyctsCgwlzDFkKD3HiVb39dU/7K/mXugcAAP//av1XvtoAAAA=
`,
	},

	"/store/migrations/1548435439_event_wrappers.down.sql": {
		name:    "1548435439_event_wrappers.down.sql",
		local:   "store/migrations/1548435439_event_wrappers.down.sql",
		size:    69,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicvQJcQ1SCHF08nFVKCgtzohPLUvNK1FwCfIPUHD29wn19VPIL0pPzMus
SizJzM+Lz0yx5uJy9vf19Qyx5gIEAAD//5tdm8pFAAAA
`,
	},

	"/store/migrations/1548435439_event_wrappers.up.sql": {
		name:    "1548435439_event_wrappers.up.sql",
		local:   "store/migrations/1548435439_event_wrappers.up.sql",
		size:    82,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/wTAwQoCIRAG4Ps8xf8entxdCWEcIcazCEl50agpoqffbwuXKI7Is4Yr1G8c8Py8H7V/
+zT448CeuSTBet3bHP9mY806brD+M0hWSGF2RHtOKaqjMwAA//8FYkZlUgAAAA==
`,
	},

	"/store/migrations/1550864142_remove_merge_field.down.sql": {
		name:    "1550864142_remove_merge_field.down.sql",
		local:   "store/migrations/1550864142_remove_merge_field.down.sql",
		size:    75,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicvQJcQ1SCHF08nFVKEoty0wtj08tS80rUXB0cVFw9vcJ9fVTyE0tSk9V
yCrOz0tS8PMPUfAL9fGx5uJy9vf19Qyx5gIEAAD//0AcbxZLAAAA
`,
	},

	"/store/migrations/1550864142_remove_merge_field.up.sql": {
		name:    "1550864142_remove_merge_field.up.sql",
		local:   "store/migrations/1550864142_remove_merge_field.up.sql",
		size:    61,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicvQJcQ1SCHF08nFVKEoty0wtj08tS80rUXAJ8g9QcPb3CfX1U8hNLUpP
tebicvb39fUMseYCBAAA//+Y8F+aPQAAAA==
`,
	},

	"/store/migrations/lock.json": {
		name:    "lock.json",
		local:   "store/migrations/lock.json",
		size:    8563,
		modtime: 1,
		compressed: `
H4sIAAAAAAAC/+xZzW7bMAy+5ykEn/sEue44IBiG7jQMBm0zLgeJ8iQqm1Pk3Ye4bWY7drYVWyunugSG
CYbfR/HP4v1KqewWCo0+W6vPK6WUuu9+lco2YDBbq6y0xiBLdvMkeGd1MPxLo6810KTqpNS9v22b7n0I
Y8kHRwZc+x7bbK3EBRxIP+IWHXJ5VOag9UC4sbIJWk/pfWL6Fo5KW9AeT5LDzWXYpUMQrHKQafhCBr2A
aWR/gUVn8zVphKa6BhoOd4Tfc9whS/7MiLpMom/2KR/Glgd/fsqBxyDviQ4z3jhH8Ex3bEnjzHniD4n5
IDXxDHJiwRpdzODPnbsUt5eWt1R19ibxF1QTR80AGHS7H8dHNP5/fPqy6nE5a6LW1cC0ByHLb6eTNs7u
qIr25H6D/liWHIOe7TqLSP16yWnTBH83asBXnzReQIJfZsSlhE9fLv+ibhlD4pc6r1TkhbiUfOE8HtpH
cMOhZUDiq7dcxMyhAI8LhX6HUC0Uen/ajbeY/tEEMnkJkGaQ1MX/IwGfG3Q1dvdQ073DWo3AMZPwNrhy
qaU3db7U+dLu4Gp2BwKuxtdZHjyYfrntwd/MNCNsaa+XcjNdFb2VIdNhYz2Jde0shfhvKDiYYi6CIkA/
LMar49PhZwAAAP//f8ci+3MhAAA=
`,
	},

	"/store/migrations": {
		name:  "migrations",
		local: `store/migrations`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"store/migrations": {
		_escData["/store/migrations/1533636853_initial.down.sql"],
		_escData["/store/migrations/1533636853_initial.up.sql"],
		_escData["/store/migrations/1537263282_add_review_target.down.sql"],
		_escData["/store/migrations/1537263282_add_review_target.up.sql"],
		_escData["/store/migrations/1537263364_add_analyzer_name_to_comment.down.sql"],
		_escData["/store/migrations/1537263364_add_analyzer_name_to_comment.up.sql"],
		_escData["/store/migrations/1537266831_add_review_old_id.down.sql"],
		_escData["/store/migrations/1537266831_add_review_old_id.up.sql"],
		_escData["/store/migrations/1537267179_add_timestamps_to_comments_and_target.down.sql"],
		_escData["/store/migrations/1537267179_add_timestamps_to_comments_and_target.up.sql"],
		_escData["/store/migrations/1537268276_data_migrate.down.sql"],
		_escData["/store/migrations/1537268276_data_migrate.up.sql"],
		_escData["/store/migrations/1537455097_delete_old_columns.down.sql"],
		_escData["/store/migrations/1537455097_delete_old_columns.up.sql"],
		_escData["/store/migrations/1537528734_review_internal_id.down.sql"],
		_escData["/store/migrations/1537528734_review_internal_id.up.sql"],
		_escData["/store/migrations/1537528735_migrate_review_event_id.down.sql"],
		_escData["/store/migrations/1537528735_migrate_review_event_id.up.sql"],
		_escData["/store/migrations/1539095669_delete_old_internal_id.down.sql"],
		_escData["/store/migrations/1539095669_delete_old_internal_id.up.sql"],
		_escData["/store/migrations/1547744983_organizations.down.sql"],
		_escData["/store/migrations/1547744983_organizations.up.sql"],
		_escData["/store/migrations/1548157581_organizations_constraints.down.sql"],
		_escData["/store/migrations/1548157581_organizations_constraints.up.sql"],
		_escData["/store/migrations/1548435439_event_wrappers.down.sql"],
		_escData["/store/migrations/1548435439_event_wrappers.up.sql"],
		_escData["/store/migrations/1550864142_remove_merge_field.down.sql"],
		_escData["/store/migrations/1550864142_remove_merge_field.up.sql"],
		_escData["/store/migrations/lock.json"],
	},
}
