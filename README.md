# go-plugin
使用情境為:

兩組件為獨立元件，並且會使用相同自訂型別，利用function call 的方式溝通並且取得自訂型別的值


兩種方式:

1. 使用go mod -trimpath 編譯兩獨立專案， 參數(-trimpath)將路徑去除並且要保證引用的package版本需一至(會去比對go.sum)的差異，需克服各專案的引用全一致

2. 使用vendor 的方式將專案是為獨立環境，引用的package視為獨立版本繞過(使用go moudle 的編譯方式)檢查，再利用unsafe.Pointer的方式取值，須確保自訂型別的正確性


GenSofile proj:
    make build-mod => 產生so檔案 (使用go mod)
    make build-mod-vendor => 產生so檔案 (使用go mod vendor)


LoadSofile 

    go-module proj (需要確保與GenSofile 的go-plugin/Iface版本一致):
        將(GenSofile proj 產生的so file 放置專案中)
        make run-mod => 執行
    
    go-vendor proj:
        將(GenSofile proj 產生的so file 放置專案中)
        make run-mod-vendor => 執行