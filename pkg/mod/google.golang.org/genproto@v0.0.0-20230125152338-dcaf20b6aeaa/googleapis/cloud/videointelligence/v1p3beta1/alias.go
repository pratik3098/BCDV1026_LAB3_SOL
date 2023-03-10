// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package videointelligence aliases all exported identifiers in package
// "cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb".
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package videointelligence

import (
	src "cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
const (
	Feature_CELEBRITY_RECOGNITION                         = src.Feature_CELEBRITY_RECOGNITION
	Feature_EXPLICIT_CONTENT_DETECTION                    = src.Feature_EXPLICIT_CONTENT_DETECTION
	Feature_FACE_DETECTION                                = src.Feature_FACE_DETECTION
	Feature_FEATURE_UNSPECIFIED                           = src.Feature_FEATURE_UNSPECIFIED
	Feature_LABEL_DETECTION                               = src.Feature_LABEL_DETECTION
	Feature_LOGO_RECOGNITION                              = src.Feature_LOGO_RECOGNITION
	Feature_OBJECT_TRACKING                               = src.Feature_OBJECT_TRACKING
	Feature_PERSON_DETECTION                              = src.Feature_PERSON_DETECTION
	Feature_SHOT_CHANGE_DETECTION                         = src.Feature_SHOT_CHANGE_DETECTION
	Feature_SPEECH_TRANSCRIPTION                          = src.Feature_SPEECH_TRANSCRIPTION
	Feature_TEXT_DETECTION                                = src.Feature_TEXT_DETECTION
	LabelDetectionMode_FRAME_MODE                         = src.LabelDetectionMode_FRAME_MODE
	LabelDetectionMode_LABEL_DETECTION_MODE_UNSPECIFIED   = src.LabelDetectionMode_LABEL_DETECTION_MODE_UNSPECIFIED
	LabelDetectionMode_SHOT_AND_FRAME_MODE                = src.LabelDetectionMode_SHOT_AND_FRAME_MODE
	LabelDetectionMode_SHOT_MODE                          = src.LabelDetectionMode_SHOT_MODE
	Likelihood_LIKELIHOOD_UNSPECIFIED                     = src.Likelihood_LIKELIHOOD_UNSPECIFIED
	Likelihood_LIKELY                                     = src.Likelihood_LIKELY
	Likelihood_POSSIBLE                                   = src.Likelihood_POSSIBLE
	Likelihood_UNLIKELY                                   = src.Likelihood_UNLIKELY
	Likelihood_VERY_LIKELY                                = src.Likelihood_VERY_LIKELY
	Likelihood_VERY_UNLIKELY                              = src.Likelihood_VERY_UNLIKELY
	StreamingFeature_STREAMING_AUTOML_ACTION_RECOGNITION  = src.StreamingFeature_STREAMING_AUTOML_ACTION_RECOGNITION
	StreamingFeature_STREAMING_AUTOML_CLASSIFICATION      = src.StreamingFeature_STREAMING_AUTOML_CLASSIFICATION
	StreamingFeature_STREAMING_AUTOML_OBJECT_TRACKING     = src.StreamingFeature_STREAMING_AUTOML_OBJECT_TRACKING
	StreamingFeature_STREAMING_EXPLICIT_CONTENT_DETECTION = src.StreamingFeature_STREAMING_EXPLICIT_CONTENT_DETECTION
	StreamingFeature_STREAMING_FEATURE_UNSPECIFIED        = src.StreamingFeature_STREAMING_FEATURE_UNSPECIFIED
	StreamingFeature_STREAMING_LABEL_DETECTION            = src.StreamingFeature_STREAMING_LABEL_DETECTION
	StreamingFeature_STREAMING_OBJECT_TRACKING            = src.StreamingFeature_STREAMING_OBJECT_TRACKING
	StreamingFeature_STREAMING_SHOT_CHANGE_DETECTION      = src.StreamingFeature_STREAMING_SHOT_CHANGE_DETECTION
)

// Deprecated: Please use vars in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
var (
	Feature_name                                                           = src.Feature_name
	Feature_value                                                          = src.Feature_value
	File_google_cloud_videointelligence_v1p3beta1_video_intelligence_proto = src.File_google_cloud_videointelligence_v1p3beta1_video_intelligence_proto
	LabelDetectionMode_name                                                = src.LabelDetectionMode_name
	LabelDetectionMode_value                                               = src.LabelDetectionMode_value
	Likelihood_name                                                        = src.Likelihood_name
	Likelihood_value                                                       = src.Likelihood_value
	StreamingFeature_name                                                  = src.StreamingFeature_name
	StreamingFeature_value                                                 = src.StreamingFeature_value
)

// Video annotation progress. Included in the `metadata` field of the
// `Operation` returned by the `GetOperation` call of the
// `google::longrunning::Operations` service.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type AnnotateVideoProgress = src.AnnotateVideoProgress

// Video annotation request.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type AnnotateVideoRequest = src.AnnotateVideoRequest

// Video annotation response. Included in the `response` field of the
// `Operation` returned by the `GetOperation` call of the
// `google::longrunning::Operations` service.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type AnnotateVideoResponse = src.AnnotateVideoResponse

// Celebrity definition.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type Celebrity = src.Celebrity

// Celebrity recognition annotation per video.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type CelebrityRecognitionAnnotation = src.CelebrityRecognitionAnnotation

// The annotation result of a celebrity face track. RecognizedCelebrity field
// could be empty if the face track does not have any matched celebrities.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type CelebrityTrack = src.CelebrityTrack

// The recognized celebrity with confidence score.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type CelebrityTrack_RecognizedCelebrity = src.CelebrityTrack_RecognizedCelebrity

// A generic detected attribute represented by name in string format.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type DetectedAttribute = src.DetectedAttribute

// A generic detected landmark represented by name in string format and a 2D
// location.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type DetectedLandmark = src.DetectedLandmark

// Detected entity from video analysis.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type Entity = src.Entity

// Explicit content annotation (based on per-frame visual signals only). If no
// explicit content has been detected in a frame, no annotations are present
// for that frame.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type ExplicitContentAnnotation = src.ExplicitContentAnnotation

// Config for EXPLICIT_CONTENT_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type ExplicitContentDetectionConfig = src.ExplicitContentDetectionConfig

// Video frame level annotation results for explicit content.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type ExplicitContentFrame = src.ExplicitContentFrame

// Face detection annotation.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type FaceDetectionAnnotation = src.FaceDetectionAnnotation

// Config for FACE_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type FaceDetectionConfig = src.FaceDetectionConfig

// Video annotation feature.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type Feature = src.Feature

// Label annotation.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type LabelAnnotation = src.LabelAnnotation

// Config for LABEL_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type LabelDetectionConfig = src.LabelDetectionConfig

// Label detection mode.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type LabelDetectionMode = src.LabelDetectionMode

// Video frame level annotation results for label detection.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type LabelFrame = src.LabelFrame

// Video segment level annotation results for label detection.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type LabelSegment = src.LabelSegment

// Bucketized representation of likelihood.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type Likelihood = src.Likelihood

// Annotation corresponding to one detected, tracked and recognized logo
// class.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type LogoRecognitionAnnotation = src.LogoRecognitionAnnotation

// Normalized bounding box. The normalized vertex coordinates are relative to
// the original image. Range: [0, 1].
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type NormalizedBoundingBox = src.NormalizedBoundingBox

// Normalized bounding polygon for text (that might not be aligned with axis).
// Contains list of the corner points in clockwise order starting from top-left
// corner. For example, for a rectangular bounding box: When the text is
// horizontal it might look like: 0----1 | | 3----2 When it's clockwise rotated
// 180 degrees around the top-left corner it becomes: 2----3 | | 1----0 and the
// vertex order will still be (0, 1, 2, 3). Note that values can be less than
// 0, or greater than 1 due to trignometric calculations for location of the
// box.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type NormalizedBoundingPoly = src.NormalizedBoundingPoly

// A vertex represents a 2D point in the image. NOTE: the normalized vertex
// coordinates are relative to the original image and range from 0 to 1.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type NormalizedVertex = src.NormalizedVertex

// Annotations corresponding to one tracked object.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type ObjectTrackingAnnotation = src.ObjectTrackingAnnotation
type ObjectTrackingAnnotation_Segment = src.ObjectTrackingAnnotation_Segment
type ObjectTrackingAnnotation_TrackId = src.ObjectTrackingAnnotation_TrackId

// Config for OBJECT_TRACKING.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type ObjectTrackingConfig = src.ObjectTrackingConfig

// Video frame level annotations for object detection and tracking. This field
// stores per frame location, time offset, and confidence.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type ObjectTrackingFrame = src.ObjectTrackingFrame

// Person detection annotation per video.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type PersonDetectionAnnotation = src.PersonDetectionAnnotation

// Config for PERSON_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type PersonDetectionConfig = src.PersonDetectionConfig

// Config for SHOT_CHANGE_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type ShotChangeDetectionConfig = src.ShotChangeDetectionConfig

// Provides "hints" to the speech recognizer to favor specific words and
// phrases in the results.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type SpeechContext = src.SpeechContext

// Alternative hypotheses (a.k.a. n-best list).
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type SpeechRecognitionAlternative = src.SpeechRecognitionAlternative

// A speech recognition result corresponding to a portion of the audio.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type SpeechTranscription = src.SpeechTranscription

// Config for SPEECH_TRANSCRIPTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type SpeechTranscriptionConfig = src.SpeechTranscriptionConfig

// The top-level message sent by the client for the `StreamingAnnotateVideo`
// method. Multiple `StreamingAnnotateVideoRequest` messages are sent. The
// first message must only contain a `StreamingVideoConfig` message. All
// subsequent messages must only contain `input_content` data.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingAnnotateVideoRequest = src.StreamingAnnotateVideoRequest
type StreamingAnnotateVideoRequest_InputContent = src.StreamingAnnotateVideoRequest_InputContent
type StreamingAnnotateVideoRequest_VideoConfig = src.StreamingAnnotateVideoRequest_VideoConfig

// `StreamingAnnotateVideoResponse` is the only message returned to the client
// by `StreamingAnnotateVideo`. A series of zero or more
// `StreamingAnnotateVideoResponse` messages are streamed back to the client.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingAnnotateVideoResponse = src.StreamingAnnotateVideoResponse

// Config for STREAMING_AUTOML_ACTION_RECOGNITION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingAutomlActionRecognitionConfig = src.StreamingAutomlActionRecognitionConfig

// Config for STREAMING_AUTOML_CLASSIFICATION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingAutomlClassificationConfig = src.StreamingAutomlClassificationConfig

// Config for STREAMING_AUTOML_OBJECT_TRACKING.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingAutomlObjectTrackingConfig = src.StreamingAutomlObjectTrackingConfig

// Config for STREAMING_EXPLICIT_CONTENT_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingExplicitContentDetectionConfig = src.StreamingExplicitContentDetectionConfig

// Streaming video annotation feature.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingFeature = src.StreamingFeature

// Config for STREAMING_LABEL_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingLabelDetectionConfig = src.StreamingLabelDetectionConfig

// Config for STREAMING_OBJECT_TRACKING.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingObjectTrackingConfig = src.StreamingObjectTrackingConfig

// Config for STREAMING_SHOT_CHANGE_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingShotChangeDetectionConfig = src.StreamingShotChangeDetectionConfig

// Config for streaming storage option.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingStorageConfig = src.StreamingStorageConfig

// Streaming annotation results corresponding to a portion of the video that
// is currently being processed.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingVideoAnnotationResults = src.StreamingVideoAnnotationResults

// Provides information to the annotator that specifies how to process the
// request.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingVideoConfig = src.StreamingVideoConfig
type StreamingVideoConfig_AutomlActionRecognitionConfig = src.StreamingVideoConfig_AutomlActionRecognitionConfig
type StreamingVideoConfig_AutomlClassificationConfig = src.StreamingVideoConfig_AutomlClassificationConfig
type StreamingVideoConfig_AutomlObjectTrackingConfig = src.StreamingVideoConfig_AutomlObjectTrackingConfig
type StreamingVideoConfig_ExplicitContentDetectionConfig = src.StreamingVideoConfig_ExplicitContentDetectionConfig
type StreamingVideoConfig_LabelDetectionConfig = src.StreamingVideoConfig_LabelDetectionConfig
type StreamingVideoConfig_ObjectTrackingConfig = src.StreamingVideoConfig_ObjectTrackingConfig
type StreamingVideoConfig_ShotChangeDetectionConfig = src.StreamingVideoConfig_ShotChangeDetectionConfig

// StreamingVideoIntelligenceServiceClient is the client API for
// StreamingVideoIntelligenceService service. For semantics around ctx use and
// closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingVideoIntelligenceServiceClient = src.StreamingVideoIntelligenceServiceClient

// StreamingVideoIntelligenceServiceServer is the server API for
// StreamingVideoIntelligenceService service.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type StreamingVideoIntelligenceServiceServer = src.StreamingVideoIntelligenceServiceServer
type StreamingVideoIntelligenceService_StreamingAnnotateVideoClient = src.StreamingVideoIntelligenceService_StreamingAnnotateVideoClient
type StreamingVideoIntelligenceService_StreamingAnnotateVideoServer = src.StreamingVideoIntelligenceService_StreamingAnnotateVideoServer

// Annotations related to one detected OCR text snippet. This will contain the
// corresponding text, confidence value, and frame level information for each
// detection.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type TextAnnotation = src.TextAnnotation

// Config for TEXT_DETECTION.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type TextDetectionConfig = src.TextDetectionConfig

// Video frame level annotation results for text annotation (OCR). Contains
// information regarding timestamp and bounding box locations for the frames
// containing detected OCR text snippets.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type TextFrame = src.TextFrame

// Video segment level annotation results for text detection.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type TextSegment = src.TextSegment

// For tracking related features. An object at time_offset with attributes,
// and located with normalized_bounding_box.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type TimestampedObject = src.TimestampedObject

// A track of an object instance.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type Track = src.Track

// UnimplementedStreamingVideoIntelligenceServiceServer can be embedded to
// have forward compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type UnimplementedStreamingVideoIntelligenceServiceServer = src.UnimplementedStreamingVideoIntelligenceServiceServer

// UnimplementedVideoIntelligenceServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type UnimplementedVideoIntelligenceServiceServer = src.UnimplementedVideoIntelligenceServiceServer

// Annotation progress for a single video.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type VideoAnnotationProgress = src.VideoAnnotationProgress

// Annotation results for a single video.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type VideoAnnotationResults = src.VideoAnnotationResults

// Video context and/or feature-specific parameters.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type VideoContext = src.VideoContext

// VideoIntelligenceServiceClient is the client API for
// VideoIntelligenceService service. For semantics around ctx use and
// closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type VideoIntelligenceServiceClient = src.VideoIntelligenceServiceClient

// VideoIntelligenceServiceServer is the server API for
// VideoIntelligenceService service.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type VideoIntelligenceServiceServer = src.VideoIntelligenceServiceServer

// Video segment.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type VideoSegment = src.VideoSegment

// Word-specific information for recognized words. Word information is only
// included in the response when certain request parameters are set, such as
// `enable_word_time_offsets`.
//
// Deprecated: Please use types in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
type WordInfo = src.WordInfo

// Deprecated: Please use funcs in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
func NewStreamingVideoIntelligenceServiceClient(cc grpc.ClientConnInterface) StreamingVideoIntelligenceServiceClient {
	return src.NewStreamingVideoIntelligenceServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
func NewVideoIntelligenceServiceClient(cc grpc.ClientConnInterface) VideoIntelligenceServiceClient {
	return src.NewVideoIntelligenceServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
func RegisterStreamingVideoIntelligenceServiceServer(s *grpc.Server, srv StreamingVideoIntelligenceServiceServer) {
	src.RegisterStreamingVideoIntelligenceServiceServer(s, srv)
}

// Deprecated: Please use funcs in: cloud.google.com/go/videointelligence/apiv1p3beta1/videointelligencepb
func RegisterVideoIntelligenceServiceServer(s *grpc.Server, srv VideoIntelligenceServiceServer) {
	src.RegisterVideoIntelligenceServiceServer(s, srv)
}
