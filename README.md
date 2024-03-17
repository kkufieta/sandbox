# sandbox

Expected Properties for:

Bunny:
```go
InternalStateFileType: ".bunny",
SensorInfo: []slam.SensorInfo{
    {Name: "bunny-camera-1", Type: slam.SensorTypeCamera},
    {Name: "bunny-imu", Type: slam.SensorTypeMovementSensor},
    {Name: "bunny-odometer", Type: slam.SensorTypeMovementSensor},
    {Name: "bunny-camera-2", Type: slam.SensorTypeCamera},
}
```

Robots:
```go
InternalStateFileType: ".robots",
SensorInfo: []slam.SensorInfo{
    {Name: "robot-movement-sensor", Type: slam.SensorTypeMovementSensor},
}
```

Fake slam:
```go
SensorInfo:  []slam.SensorInfo{}
```