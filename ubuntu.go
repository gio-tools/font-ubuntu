package ubuntu

import (
	"sync"

	"gio.tools/fonts/ubuntu/ubuntubold"
	"gio.tools/fonts/ubuntu/ubuntubolditalic"
	"gio.tools/fonts/ubuntu/ubuntuitalic"
	"gio.tools/fonts/ubuntu/ubuntulight"
	"gio.tools/fonts/ubuntu/ubuntulightitalic"
	"gio.tools/fonts/ubuntu/ubuntumedium"
	"gio.tools/fonts/ubuntu/ubuntumediumitalic"
	"gio.tools/fonts/ubuntu/ubunturegular"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

var (
	once       sync.Once
	collection []font.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(ubuntubold.TTF)
		register(ubuntubolditalic.TTF)
		register(ubuntuitalic.TTF)
		register(ubuntulight.TTF)
		register(ubuntulightitalic.TTF)
		register(ubuntumedium.TTF)
		register(ubuntumediumitalic.TTF)
		register(ubunturegular.TTF)
		// Ensure that any outside appends will not reuse the backing store.
		n := len(collection)
		collection = collection[:n:n]
	})
	return collection
}

func register(src []byte) {
	faces, err := opentype.ParseCollection(src)
	if err != nil {
		panic("failed to parse font: " + err.Error())
	}
	collection = append(collection, faces[0])
}
