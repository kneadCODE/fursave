//go:generate sqlboiler --wipe -c sqlboiler.yaml psql

package ledgersvc

import (
	_ "github.com/friendsofgo/errors"                         // for sqlboiler.
	_ "github.com/volatiletech/null/v8"                       // for sqlboiler.
	_ "github.com/volatiletech/sqlboiler/v4/boil"             // for sqlboiler.
	_ "github.com/volatiletech/sqlboiler/v4/drivers"          // for sqlboiler.
	_ "github.com/volatiletech/sqlboiler/v4/queries"          // for sqlboiler.
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"       // for sqlboiler.
	_ "github.com/volatiletech/sqlboiler/v4/queries/qmhelper" // for sqlboiler.
	_ "github.com/volatiletech/sqlboiler/v4/types"            // for sqlboiler.
	_ "github.com/volatiletech/strmangle"                     // for sqlboiler.
)
