package io

interface Interfacer {
  Input(i interface{})
  Output(i interface{})
}

struct Interface {
  parser Parser
  game BoardGame
}
