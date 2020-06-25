// Sample speech-quickstart uses the Google Cloud Speech API to transcribe
// audio.
package main

import (
        "context"
        "fmt"
        "io/ioutil"
        "log"

        speech "cloud.google.com/go/speech/apiv1"
        speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func main() {
        ctx := context.Background()

        // Creates a client.
        client, err := speech.NewClient(ctx)
        if err != nil {
                log.Fatalf("Failed to create client: %v", err)
        }

        // ファイルを変更したい場合はココを書き換える
        filename := "sample.wav"

        // Reads the audio file into memory.
        data, err := ioutil.ReadFile(filename)
        if err != nil {
                log.Fatalf("Failed to read file: %v", err)
        }

        // Detects speech in the audio file.
        resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
                Config: &speechpb.RecognitionConfig{
                        // wavファイルではヘッダに入っているのでコメントアウト
                        //Encoding:        speechpb.RecognitionConfig_LINEAR16,
                        // wavファイルではヘッダに入っているのでコメントアウト
                        //SampleRateHertz: 16000,
                        // 必須パラメーた
                        LanguageCode:    "en-US",
                },
                Audio: &speechpb.RecognitionAudio{
                        AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
                },
        })
        if err != nil {
                log.Fatalf("failed to recognize: %v", err)
        }

        // Prints the results.
        for _, result := range resp.Results {
                for _, alt := range result.Alternatives {
                        fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
                }
        }
}
