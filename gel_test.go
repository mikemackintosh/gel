package gel

import (
  "fmt"
  "os"
  "testing"

)

var TEST_ENV_KEY = "TEST_KEY_GEL"

var useCaseEnvFlagsConfig = []struct{
  Env string
  Cli string
  Want string
  Expect bool
} {
  {
    Env: "from_env",
    Cli: "from_cli",
    Want: "from_env",
    Expect: true,
  },
  {
    Env: "234234234234234",
    Cli: "asdffadsfdsafas",
    Want: "from_env",
    Expect: false,
  },
}

func TestUseOrderEnvFlagsConfig(t *testing.T) {
  for _, test := range useCaseEnvFlagsConfig {
    // Start with a fresh registry for every iteration
    registry = NewRegistry()

    // Set tje test values
    os.Setenv(TEST_ENV_KEY, test.Env)
    os.Args = []string{fmt.Sprintf("-%s=%s", TEST_ENV_KEY, test.Cli)}

    // Set the var
    String(TEST_ENV_KEY, "default", "test value")

    // Set UseOrder
    UseOrder(Env, Flags)

    // Load Up!
    Up()

    v := MustGet(TEST_ENV_KEY).String()
    if test.Expect == (v != test.Want) {
      fmt.Printf("\tWanted '%s', got '%s'\n", test.Want, v)
      t.Fail()
    }
  }
}


var useCaseFlagsConfigEnv = []struct{
  Env string
  Cli string
  Want string
  Expect bool
} {
  {
    Env: "from_env",
    Cli: "from_cli",
    Want: "from_cli",
    Expect: true,
  },
  {
    Env: "234234234234234",
    Cli: "asdffadsfdsafas",
    Want: "from_cli",
    Expect: false,
  },
}

func TestUseOrderFlagsConfigEnv (t *testing.T) {
  for _, test := range useCaseFlagsConfigEnv {
    // Start with a fresh registry for every iteration
    registry = NewRegistry()

    // Set tje test values
    os.Setenv(TEST_ENV_KEY, test.Env)
    os.Args = []string{"test", fmt.Sprintf("-%s=%s", TEST_ENV_KEY, test.Cli)}

    // Set the var
    String(TEST_ENV_KEY, "default", "test value")

    // Set UseOrder
    UseOrder(Flags, Env)

    // Load Up!
    Up()

    v := MustGet(TEST_ENV_KEY).String()
    if test.Expect == (v != test.Want) {
      fmt.Printf("\tWanted '%s', got '%s'\n", test.Want, v)
      t.Fail()
    }
  }
}
