package play

type Thrust struct {
	velocity     float64
	maxSpeed     float64
	acceleration float64
	deceleration float64
}

func NewThrust(maxSpeed, acceleration, deceleration float64) *Thrust {
	return &Thrust{
		velocity:     0,
		maxSpeed:     maxSpeed,
		acceleration: acceleration,
		deceleration: deceleration,
	}
}

func (t *Thrust) Advance(delta uint64, up, down bool) int32 {
	deltaSeconds := float64(delta) / 1000.0

	// apply input forces - immediate direction change allowed;
	// disable controls if destroyed
	if up {
		t.velocity -= t.acceleration * deltaSeconds
	} else if down {
		t.velocity += t.acceleration * deltaSeconds
	} else {
		// no input - decelerate towards zero
		if t.velocity > 0 {
			t.velocity -= t.deceleration * deltaSeconds

			if t.velocity < 0 {
				t.velocity = 0
			}
		} else if t.velocity < 0 {
			t.velocity += t.deceleration * deltaSeconds

			if t.velocity > 0 {
				t.velocity = 0
			}
		}
	}

	// clamp velocity to max speed
	if t.velocity > t.maxSpeed {
		t.velocity = t.maxSpeed
	} else if t.velocity < -t.maxSpeed {
		t.velocity = -t.maxSpeed
	}

	return int32(t.velocity * deltaSeconds)
}

func (t *Thrust) FullStop() {
	t.velocity = 0
}

func (t *Thrust) Percentage() int {
	return int(t.velocity / t.maxSpeed * 100)
}
