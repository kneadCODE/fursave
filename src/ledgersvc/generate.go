//go:generate sqlboiler --wipe -c sqlboiler.yaml psql

package ledgersvc

import (
	// sqlboiler dependencies:-
	_ "github.com/friendsofgo/errors"
	_ "github.com/volatiletech/null/v8"
	_ "github.com/volatiletech/sqlboiler/v4/boil"
	_ "github.com/volatiletech/sqlboiler/v4/drivers"
	_ "github.com/volatiletech/sqlboiler/v4/queries"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	_ "github.com/volatiletech/sqlboiler/v4/types"
	_ "github.com/volatiletech/strmangle"
)
