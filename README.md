# KrunkerAPI
Krunker api with GoLang

## Player Profile Request Map

### Main Fields

- `player_name`: Player's name.
- `player_clan`: Clan name.
- `player_kills`: Total number of kills by the player.
- `player_deaths`: Total number of deaths.
- `player_score`: Player's total score.
- `player_timeplayed`: Total playtime in milliseconds.
- `player_games_played`: Total number of games played.
- `player_wins`: Total games won.
- `player_funds`: KR (in-game currency) owned by the player.
- `player_skinvalue`: Total value of owned skins.
- `player_datenew`: Date the player registered.
- `player_followed`: Number of followers.
- `player_following`: Number of people the player is following.
- `player_elo4`: Junk data or unused.

### Unclear Fields

- `clan_rank`: Possibly the rank within the clan.
- `partner_approved`: Related to verified badges.
- `player_alias`: Premium alias, if applicable.
- `player_badge`: Represents player's badges.
- `player_calls`: Possibly the number of KPD (Krunker Police Department) reports, but unclear.
- `player_chal`: Might relate to KPD reports, unclear.
- `player_elo`: Unknown usage.
- `player_elo2`: Unknown usage.
- `player_eventcount`: Unclear purpose.
- `player_featured`: Unknown usage.
- `player_hack`: Possibly flags a player for hacking but unclear.
- `player_id`: Player ID, unclear usage.
- `player_infected`: Possibly related to infection in zombie mode.
- `player_jobrating`: Player's rating in the game, unclear context.
- `player_jobratingpositive`: Positive job ratings, unclear context.
- `player_premium`: Premium badge status.
- `player_region`: Region or server location.
- `player_taggedaccounts`: Number of tagged accounts, unclear.
- `player_twitchname`: Linked Twitch account name, if any.
- `player_type`: Unclear usage.

### Player Stats

- `n`: Number of nukes.
- `s`: Total shots fired.
- `h`: Total hits.
- `hs`: Total headshots.
- `ls`: Total legshots.
- `wb`: Total wall bangs.
- `mk`: Total melee kills.
- `tmk`: Total throwing melee kills.
- `fk`: Total fist kills.
- `spry`: Total sprays used.
- `crc`: Total crouches.
- `sl`: Total slimers.

### Unknown `player_stats` Keys

- `c`, `c1`, `c2`, `ast`, `c5`, `c8`, `r2`, `c0`, `c4`, `c12`, `c11`, `r4`, `r3`, `c3`, `c6`, `flg`, `c7`, `c9`, `c15`, `sad`, `c13`, `r1`, `bdg` (likely matches `player_badge`), `ad`, `cad`, `c10`, `c14`, `r5`: Unclear or unknown usages.

---

## API Usage Example

Initialize using `NewKrunkerAPI`.  
Return value: `KrunkerAPI`, `error`.

To reserve structure destruction, use `defer api.Close()`.

### Fetching Player Profile

`GetProfile(username string)`

Return value: `*Profile`, `message`.  
This returns a pointer to the `Profile` object along with the original `decodedMessage`.
If the profile does not exist, it returns nill.


The keys and values from `decodedMessage` are outlined above, but some fields remain unclear. Help would be appreciated for those.

### Example Code

```go
package main

import (
    KrunkerAPI "krunker-api"
    "log"
)

func main() {
    api, _ := KrunkerAPI.NewKrunkerAPI()
    profile, _ := api.GetProfile("a6a6")

    log.Println(*profile)
}
```

# **Do not use for commercial purposes**
This project is under development, if you see any bugs or areas that need fixing, please create an issue!
