{{$args := parseArgs 2 "Custom usage text when invalid args provided"
    (carg "string" "event")
	(carg "int" "players")
	(carg "string" "desc")
}}

{{ $eventName := $args.Get 0}}
{{ $requiredPlayers := $args.Get 1}}
{{ $desc := $args.Get 2}}
{{ /* key used for the player, usually restricted to one key: "lfg" */ }}
{{ $key := "lfg" }}

{{ /* using placeholders here because the test server doesn't have the tank/dps/heal emojis */ }}
{{ $tankEmoji := ":ok_hand:" }} {{ $healEmoji := ":100:" }} {{ $dpsEmoji := ":eggplant:"}}


{{ $newLfgMessageString := joinStr "" .Member.User "is looking for " $requiredPlayers " for " $eventName "!"}}
{{ $newLfgMessage := sendMessageRetID nil $newLfgMessageString }}
{{ addMessageReactions nil $newLfgMessage $tankEmoji $healEmoji $dpsEmoji }}
{{ dbSet .Member.User.ID $key $newLfgMessage}}