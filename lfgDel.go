{{ /* key used for the player, usually restricted to one key: "lfg" */ }}
{{ $key := "lfg" }}
{{ $message := joinStr "" "Deleted: " (dbGet .Member.User.ID $key).Value }
{{ dbDel .Member.User.ID $key }}
{{ sendMessageRetID nil $message }}