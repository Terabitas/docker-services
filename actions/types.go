package actions

type (
    Setuper interface {
        Setup() error
    }

    Updater interface {
        Update() error
    }

    Remover interface {
        Remove() error
    }

    Starter interface {
        Start() error
    }

    Stopper interface {
        Stop() error
    }

    Describer interface {
        Describe() error
    }
)
