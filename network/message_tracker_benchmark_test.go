package network

import (
 "testing"

 "github.com/ChainSafe/gossamer-go-interview/network"
)

func BenchmarkAdd(b *testing.B) {
 tracker := network.NewMessageTracker(1000)
 message := &network.Message{ID: "1", PeerID: "peer1", Data: []byte("data")}

 b.ReportAllocs()
 b.ResetTimer()

 for i := 0; i < b.N; i++ {
  _ = tracker.Add(message)
 }
}

func BenchmarkDelete(b *testing.B) {
 tracker := network.NewMessageTracker(1000)
 message := &network.Message{ID: "1", PeerID: "peer1", Data: []byte("data")}
 _ = tracker.Add(message)

 b.ReportAllocs()
 b.ResetTimer()

 for i := 0; i < b.N; i++ {
  _ = tracker.Delete(message.ID)
 }
}

func BenchmarkMessage(b *testing.B) {
 tracker := network.NewMessageTracker(1000)
 message := &network.Message{ID: "1", PeerID: "peer1", Data: []byte("data")}
 _ = tracker.Add(message)

 b.ReportAllocs()
 b.ResetTimer()

 for i := 0; i < b.N; i++ {
  _, _ = tracker.Message(message.ID)
 }
}

func BenchmarkMessages(b *testing.B) {
 tracker := network.NewMessageTracker(1000)
 message := &network.Message{ID: "1", PeerID: "peer1", Data: []byte("data")}
 _ = tracker.Add(message)

 b.ReportAllocs()
 b.ResetTimer()

 for i := 0; i < b.N; i++ {
  _ = tracker.Messages()
 }
}
