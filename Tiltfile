version_settings() # Tilt version settings
allow_k8s_contexts('k3d-cluster') # Specify k3d cluster

# Main configuration
def ledgersvc():
    docker_build(
        'image-serverd:dev',
        './src/ledgersvc',
        dockerfile='./src/ledgersvc/build/Dockerfile',
        # live_update=[
        #     sync('./src/ledgersvc', '/go/src/app'),
        #     restart_container()
        # ]
    )

    k8s_yaml(kustomize('./src/ledgersvc/build/k8s/overlays/development'))

    k8s_resource(
        'ledgersvc-postgres',
        port_forwards=[port_forward(90001, 5432, name='ledgersvc-postgres')],
        labels=['leggersvc-postgres']
    )
    k8s_resource(
        'ledgersvc-serverd',
        port_forwards=[port_forward(9000, 80, name='ledgersvc-serverd')],
        resource_deps=['ledgersvc-postgres'],
        labels=['leggersvc-serverd']
    )

# Trigger the configuration
ledgersvc()

# Watch settings for faster development
watch_settings(
    ignore=[
        '.git',
        '.vscode',
        '**/node_modules',
        '**/*.log',
        # '**/build',
        '**/dist'
    ]
)
