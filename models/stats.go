package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// AggregatedStats are the aggregated stats for an orchestrator
type AggregatedStats struct {
	ID          string  `json:"-" bson:"_id,omitempty"`
	Score       float64 `bson:"score" json:"score"`
	SuccessRate float64 `bson:"success_rate" json:"success_rate"`
}

// Stats are the raw stats per test stream
type Stats struct {
	ID                primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Region            string             `json:"region" bson:"-"`
	Orchestrator      string             `json:"orchestrator" bson:"orchestrator"`
	SegmentsSent      int                `json:"segments_sent" bson:"segments_sent"`
	SegmentsReceived  int                `json:"segments_received" bson:"segments_received"`
	SuccessRate       float64            `json:"success_rate" bson:"success_rate"`
	AvgSegDuration    int64              `json:"avg_seg_duration" bson:"avg_seg_duration"`
	AvgUploadTime     int64              `json:"avg_upload_time" bson:"avg_upload_time"`
	AvgUploadScore    float64            `json:"avg_upload_score" bson:"avg_upload_score"`
	AvgDownloadTime   int64              `json:"avg_download_time" bson:"avg_download_time"`
	AvgDownloadScore  float64            `json:"avg_download_score" bson:"avg_download_score"`
	AvgTranscodeTime  int64              `json:"avg_transcode_time" bson:"avg_transcode_time"`
	AvgTranscodeScore float64            `json:"avg_transcode_score" bson:"avg_transcode_score"`
	RoundTripScore    float64            `json:"round_trip_score" bson:"round_trip_score"`
	Errors            []string           `json:"errors" bson:"errors"`
	Timestamp         int64              `json:"timestamp" bson:"timestamp"`
}