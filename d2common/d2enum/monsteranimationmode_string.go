// Code generated by "stringer -linecomment -type MonsterAnimationMode"; DO NOT EDIT.

package d2enum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AnimationModeMonsterDeath-0]
	_ = x[AnimationModeMonsterNeutral-1]
	_ = x[AnimationModeMonsterWalk-2]
	_ = x[AnimationModeMonsterGetHit-3]
	_ = x[AnimationModeMonsterAttack1-4]
	_ = x[AnimationModeMonsterAttack2-5]
	_ = x[AnimationModeMonsterBlock-6]
	_ = x[AnimationModeMonsterCast-7]
	_ = x[AnimationModeMonsterSkill1-8]
	_ = x[AnimationModeMonsterSkill2-9]
	_ = x[AnimationModeMonsterSkill3-10]
	_ = x[AnimationModeMonsterSkill4-11]
	_ = x[AnimationModeMonsterDead-12]
	_ = x[AnimationModeMonsterKnockback-13]
	_ = x[AnimationModeMonsterSequence-14]
	_ = x[AnimationModeMonsterRun-15]
}

const _MonsterAnimationMode_name = "DTNUWLGHA1A2BLSCS1S2S3S4DDGHxxRN"

var _MonsterAnimationMode_index = [...]uint8{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32}

func (i MonsterAnimationMode) String() string {
	if i < 0 || i >= MonsterAnimationMode(len(_MonsterAnimationMode_index)-1) {
		return "MonsterAnimationMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MonsterAnimationMode_name[_MonsterAnimationMode_index[i]:_MonsterAnimationMode_index[i+1]]
}