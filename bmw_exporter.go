package main

import (
  "fmt"
  "net/http"
  "strings"
  "log"
  "os/exec"
)

func exec_shell(command string) string {
        out, err := exec.Command("/bin/bash", "-c", command).Output()
        if err != nil {
                log.Fatal(err)
        }
        return string(out)
}


func metricsHandler(w http.ResponseWriter, r *http.Request) {
        location_lat_raw := exec_shell("cat bmw.json  | jq '.[].location.coordinates.latitude | select( . != null)'")
        location_lat := strings.Replace(location_lat_raw, "\n", "", -1)
        fmt.Fprintf(w, "location_lat{label=\"geo\"}  %s\n" , location_lat)

        location_long_raw := exec_shell("cat bmw.json  | jq '.[].location.coordinates.longitude | select( . != null)'")
        location_long := strings.Replace(location_long_raw, "\n", "", -1)
        fmt.Fprintf(w,"location_long{label=\"geo\"} %s\n" ,location_long)

        checkControlMessages_severity_raw := exec_shell("cat bmw.json  | jq '.[].checkControlMessages | select( . != null)' | jq '.[].severity'")
        checkControlMessages_severity := strings.Replace(checkControlMessages_severity_raw, "\n", "", -1)
        fmt.Fprintf(w,"checkControlMessages_severity %d\n" ,len(checkControlMessages_severity))

        checkControlMessages_type_raw := exec_shell("cat bmw.json  | jq '.[].checkControlMessages | select( . != null)' | jq '.[].type'")
        checkControlMessages_type := strings.Replace(checkControlMessages_type_raw, "\n", "", -1)
        fmt.Fprintf(w,"checkControlMessages_type %d\n" ,len(checkControlMessages_type))

        combustionFuelLevel_range_raw := exec_shell("cat bmw.json  | jq '.[].combustionFuelLevel.range | select( . != null)'")
        combustionFuelLevel_range := strings.Replace(combustionFuelLevel_range_raw, "\n", "", -1)
        fmt.Fprintf(w,"combustionFuelLevel_range %s\n" ,combustionFuelLevel_range)

        fuelPrice:= exec_shell("curl -s \"https://wafi.iit.cnr.it/lab/djst/opencarburanti/api/getTemporalValues.php?brand=tamoil&carburante=benzina\" | jq \".\" | tail -n4 | grep value | awk '{print $2}' | awk -F '\"' '{ print$2 }'")
        fmt.Fprintf(w,"gas_fuel_price %s\n" ,fuelPrice)

        combustionFuelLevel_remainingFuelLiters_raw := exec_shell("cat bmw.json  | jq '.[].combustionFuelLevel.remainingFuelLiters | select( . != null)'")
        combustionFuelLevel_remainingFuelLiters := strings.Replace(combustionFuelLevel_remainingFuelLiters_raw, "\n", "", -1)
        fmt.Fprintf(w,"combustionFuelLevel_remainingFuelLiters %s\n" ,combustionFuelLevel_remainingFuelLiters)

        currentMileages_raw := exec_shell("cat bmw.json  | jq '.[].currentMileage | select( . != null)'")
        currentMileages := strings.Replace(currentMileages_raw, "\n", "", -1)
        fmt.Fprintf(w,"currentMileages %s\n" ,currentMileages)

        autonomy_range_raw := exec_shell("cat bmw.json  | jq '.[].range | select( . != null)'")
        autonomy_range := strings.Replace(autonomy_range_raw, "\n", "", -1)
        fmt.Fprintf(w,"range %s\n" ,autonomy_range)

        doorsState_combinedSecurityState_raw := exec_shell("cat bmw.json  | jq '.[].doorsState.combinedSecurityState | select( . != null)'")
        doorsState_combinedSecurityState := strings.Replace(doorsState_combinedSecurityState_raw, "\n", "", -1)
        fmt.Fprintf(w,"doorsState_combinedSecurityState %d\n" ,len(doorsState_combinedSecurityState))

        doorsState_combinedState_raw := exec_shell("cat bmw.json  | jq '.[].doorsState.combinedState | select( . != null)'")
        doorsState_combinedState := strings.Replace(doorsState_combinedState_raw, "\n", "", -1)
        fmt.Fprintf(w,"doorsState_combinedState %d\n" ,len(doorsState_combinedState))

        location_heading_raw := exec_shell("cat bmw.json  | jq '.[].location.heading | select( . != null)'")
        location_heading := strings.Replace(location_heading_raw, "\n", "", -1)
        fmt.Fprintf(w,"location_heading %s\n" ,location_heading)


        doorsState_hood_raw := exec_shell("cat bmw.json  | jq '.[].doorsState.hood | select( . != null)'")
        doorsState_hood := strings.Replace(doorsState_hood_raw, "\n", "", -1)
        fmt.Fprintf(w,"hood %d\n" ,len(doorsState_hood))

        doorsState_leftFront_raw := exec_shell("cat bmw.json  | jq '.[].doorsState.leftFront | select( . != null)'")
        doorsState_leftFront := strings.Replace(doorsState_leftFront_raw, "\n", "", -1)
        fmt.Fprintf(w,"leftFront %d\n" ,len(doorsState_leftFront))

        doorsState_leftRear_raw := exec_shell("cat bmw.json  | jq '.[].doorsState.leftRear | select( . != null)'")
        doorsState_leftRear := strings.Replace(doorsState_leftRear_raw, "\n", "", -1)
        fmt.Fprintf(w,"leftRear %d\n" ,len( doorsState_leftRear))

        doorsState_rightFront_raw := exec_shell("cat bmw.json  | jq '.[].doorsState.rightFront | select( . != null)'")
        doorsState_rightFront := strings.Replace(doorsState_rightFront_raw, "\n", "", -1)
        fmt.Fprintf(w,"rightFront %d\n" ,len( doorsState_rightFront))

        doorsState_rightRear_raw := exec_shell("cat bmw.json  | jq '.[].doorsState.rightRear | select( . != null)'")
        doorsState_rightRear := strings.Replace(doorsState_rightRear_raw, "\n", "", -1)
        fmt.Fprintf(w,"rightRear %d\n" ,len( doorsState_rightRear))

        doorsState_trunk_raw := exec_shell("cat bmw.json  | jq '.[].doorsState.trunk | select( . != null)'")
        doorsState_trunk := strings.Replace(doorsState_trunk_raw, "\n", "", -1)
        fmt.Fprintf(w,"trunk %d\n" ,len( doorsState_trunk))

        windowsState_combinedState_raw := exec_shell("cat bmw.json  | jq '.[].windowsState.combinedState | select( . != null)'")
        windowsState_combinedState := strings.Replace(windowsState_combinedState_raw, "\n", "", -1)
        fmt.Fprintf(w,"windowsStatecombinedState %d\n" ,len( windowsState_combinedState))

        windowsState_leftFront_raw := exec_shell("cat bmw.json  | jq '.[].windowsState.leftFront | select( . != null)'")
        windowsState_leftFront := strings.Replace(windowsState_leftFront_raw, "\n", "", -1)
        fmt.Fprintf(w,"windowsStateleftFront %d\n" ,len( windowsState_leftFront))

        windowsState_leftRear_raw := exec_shell("cat bmw.json  | jq '.[].windowsState.leftRear | select( . != null)'")
        windowsState_leftRear := strings.Replace(windowsState_leftRear_raw, "\n", "", -1)
        fmt.Fprintf(w,"windowsStateleftRear %d\n" ,len( windowsState_leftRear))

        windowsState_rightFront_raw := exec_shell("cat bmw.json  | jq '.[].windowsState.rightFront | select( . != null)'")
        windowsState_rightFront := strings.Replace(windowsState_rightFront_raw, "\n", "", -1)
        fmt.Fprintf(w,"windowsStaterightFront %d\n" ,len( windowsState_rightFront))

        windowsState_rightRear_raw := exec_shell("cat bmw.json  | jq '.[].windowsState.rightRear | select( . != null)'")
        windowsState_rightRear := strings.Replace(windowsState_rightRear_raw, "\n", "", -1)
        fmt.Fprintf(w,"windowsStaterightRear %d\n" ,len( windowsState_rightRear))

}

func main() {

        http.HandleFunc("/", metricsHandler)
        http.ListenAndServe(":1337", nil)

}