{{ /* key used for the player, usually restricted to one key: "lfg" */ }}
{{ $key := "lfg" }}
{{ sendMessageRetID nil (dbGet .Member.User.ID $key).Value }}