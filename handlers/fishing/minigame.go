package fishing

import (
	albionlens "github.com/AutoDruid/albion-lens"
	"github.com/AutoDruid/albion-lens/types/events"
	"github.com/AutoDruid/photon-parser"
)

type FishingMinigameEvent struct {
	PlayerObjectID uint64 `json:"player_object_id"` // param[0]
	RodTier        uint8  `json:"rod_tier"`         // param[1]
	ServerTicks    int64  `json:"server_ticks"`     // param[2]

	// Core decision variables
	PlayerIndicator float32 `json:"player_indicator"` // param[3] — your position in the zone
	FishIndicator   float32 `json:"fish_indicator"`   // param[4] — fish tension position
	ZoneUpperBound  float32 `json:"zone_upper_bound"` // param[5] — upper safe boundary
	ZoneLowerBound  float32 `json:"zone_lower_bound"` // param[9] — lower safe boundary (negative)

	// Physics constants (stable per fish type)
	TensionIncRate  float32 `json:"tension_inc_rate"`  // param[6]  = 0.3
	TensionDecRate  float32 `json:"tension_dec_rate"`  // param[7]  = 1.1
	CastDistance    float32 `json:"cast_distance"`     // param[8]  — varies
	TotalRangeMax   float32 `json:"total_range_max"`   // param[11] = 6.0
	ZoneVisualWidth float32 `json:"zone_visual_width"` // param[12] — safe zone render width
	TimeoutSecs     float32 `json:"timeout_secs"`      // param[13] = 10.0
	DifficultyMax   float32 `json:"difficulty_max"`    // param[15] = 7.0
}

func decodeFishingMinigame(params []photon.ParameterV18) (FishingMinigameEvent, bool) {
	var e FishingMinigameEvent

	v, ok := params[0].IntValue()
	if !ok {
		return e, false
	}
	e.PlayerObjectID = uint64(v)

	v, ok = params[1].IntValue()
	if !ok {
		return e, false
	}
	e.RodTier = uint8(v)

	v, ok = params[2].IntValue()
	if !ok {
		return e, false
	}
	e.ServerTicks = v

	playerIndicator, ok := params[3].Float32Value()
	if !ok {
		return e, false
	}
	e.PlayerIndicator = playerIndicator

	fishIndicator, ok := params[4].Float32Value()
	if !ok {
		return e, false
	}
	e.FishIndicator = fishIndicator

	zoneUpperBound, ok := params[5].Float32Value()
	if !ok {
		return e, false
	}
	e.ZoneUpperBound = zoneUpperBound

	zoneLowerBound, ok := params[9].Float32Value()
	if !ok {
		return e, false
	}
	e.ZoneLowerBound = zoneLowerBound

	tensionIncRate, ok := params[6].Float32Value()
	if !ok {
		return e, false
	}
	e.TensionIncRate = tensionIncRate

	tensionDecRate, ok := params[7].Float32Value()
	if !ok {
		return e, false
	}
	e.TensionDecRate = tensionDecRate

	castDistance, ok := params[8].Float32Value()
	if !ok {
		return e, false
	}
	e.CastDistance = castDistance

	totalRangeMax, ok := params[11].Float32Value()
	if !ok {
		return e, false
	}
	e.TotalRangeMax = totalRangeMax

	zoneVisualWidth, ok := params[12].Float32Value()
	if !ok {
		return e, false
	}
	e.ZoneVisualWidth = zoneVisualWidth

	timeoutSecs, ok := params[13].Float32Value()
	if !ok {
		return e, false
	}
	e.TimeoutSecs = timeoutSecs

	difficultyMax, ok := params[14].Float32Value()
	if !ok {
		return e, false
	}
	e.DifficultyMax = difficultyMax

	return e, true
}

var FishingMinigameHandler = albionlens.EventHandler[FishingMinigameEvent]{
	Code:    events.FishingMiniGame,
	Handler: decodeFishingMinigame,
}
