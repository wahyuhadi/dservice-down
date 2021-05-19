package main

import "testing"

func TestSpawnDial(t *testing.T) {
	t.Run("Testing with localhost", func(t *testing.T) {
		var ip string = "127.0.0.1"
		status := Ping(ip)

		if !status {
			t.Errorf("Expectation is true, but return is %t", status)
			t.Fail()
		}
	})

	t.Run("Testing with bad ip format", func(t *testing.T) {
		var ip string = "127.0.0.1232"
		status := Ping(ip)

		if status {
			t.Errorf("Expectation is false, but return is %t", status)
			t.Fail()
		}
	})

	t.Run("Testing with no live host", func(t *testing.T) {
		var ip string = "137.0.0.2"
		status := Ping(ip)

		if status {
			t.Errorf("Expectation is false, but return is %t", status)
			t.Fail()
		}
	})
}
