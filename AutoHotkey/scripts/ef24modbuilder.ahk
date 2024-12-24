#Requires AutoHotkey v2.0
#Include utils/resources.ahk

BundlePath := "C:\Program Files (x86)\Steam\steamapps\common\VTOL VR\DLC\2531290\2531290"
ResourceFile1 := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061"
ResourceFile2 := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource"
ResourceFile3 := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS"
AssetsPath := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061-mod"
ResourcePath := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resource-mod"
ResSPath := "C:\Users\Administrator\Desktop\CAB-db515831ae078197daa2fd6af388d061.resS-mod"
ExportPath:= "C:\Users\Administrator\Desktop\"


Sleep 5000
UnpackBundle(BundlePath, ExportPath)
OpenResourceFiles(ResourceFile1,ResourceFile2,ResourceFile3)
ReplaceEFResources
SaveMod(AssetsPath, ResourcePath, ResSPath)
CleanupUABE
CleanupAva

ReplaceEFResources(){
    ReplaceTextureResource("dashSprites",           "C:\resources\dashSprites",             ""          )
    ReplaceTextureResource("rwrTexture",            "C:\resources\rwrTexture",              "{F3}"      )
    ReplaceAudioResource("missileLockLoop",         "C:\resources\missileLockLoop",         "{f3}"      )
    ReplaceAudioResource("mwsTone",                 "C:\resources\mwsTone",                 ""          )
    ReplaceAudioResource("radarLockLoop",           "C:\resources\radarLockLoop",           ""          )
    ReplaceAudioResource("rwrNewContact2",          "C:\resources\rwrNewContact2",          ""          )
    ReplaceAudioResource("rwrPing2",                 "C:\resources\rwrPing2",               ""          )
    ReplaceAudioResource("SARHLockLoop",            "C:\resources\SARHLockLoop",            ""          )
    ReplaceAudioResource("ttsw_altitude",           "C:\resources\ttsw_altitude",           ""          )
    ReplaceAudioResource("ttsw_autopilotOff",       "C:\resources\ttsw_autopilotOff",       ""          )
    ReplaceAudioResource("ttsw_auxEngineFailure",   "C:\resources\ttsw_auxEngineFailure",   ""          )
    ReplaceAudioResource("ttsw_bingoFuel",          "C:\resources\ttsw_bingoFuel",          ""          )
    ReplaceAudioResource("ttsw_chaff",              "C:\resources\ttsw_chaff",              "{F3}{F3}"  )
    ReplaceAudioResource("ttsw_chaffEmpty",         "C:\resources\ttsw_chaffEmpty",         ""          )
    ReplaceAudioResource("ttsw_chaffLow",           "C:\resources\ttsw_chaffLow",           ""          )
    ReplaceAudioResource("ttsw_engineFailure",      "C:\resources\ttsw_engineFailure",      ""          )
    ReplaceAudioResource("ttsw_flare",              "C:\resources\ttsw_flare",              "{F3}"      )
    ReplaceAudioResource("ttsw_flareEmpty",         "C:\resources\ttsw_flareEmpty",         ""          )
    ReplaceAudioResource("ttsw_flareLow",           "C:\resources\ttsw_flareLow",           ""          )
    ReplaceAudioResource("ttsw_fuelDump",           "C:\resources\ttsw_fuelDump",           ""          )
    ReplaceAudioResource("ttsw_landingGear",        "C:\resources\ttsw_landingGear",        ""          )
    ReplaceAudioResource("ttsw_leftEngineFailure",  "C:\resources\ttsw_leftEngineFailure",  ""          )
    ReplaceAudioResource("ttsw_missileLaunch",      "C:\resources\ttsw_missileLaunch",      ""          )
    ReplaceAudioResource("ttsw_overG",              "C:\resources\ttsw_overG",              ""          )
    ReplaceAudioResource("ttsw_pitbull",            "C:\resources\ttsw_pitbull",            ""          )
    ReplaceAudioResource("ttsw_pullUp",             "C:\resources\ttsw_pullUp",             ""          )
    ReplaceAudioResource("ttsw_rightEngineFailure", "C:\resources\ttsw_rightEngineFailure", ""          )
    ReplaceAudioResource("ttsw_shoot",              "C:\resources\ttsw_shoot",              ""          )
    ReplaceAudioResource("ttsw_wingFold",           "C:\resources\ttsw_wingFold",           ""          )
    ReplaceAudioResource("warningBeep",             "C:\resources\warningBeep",             "{F3}{F3}"  )
}
