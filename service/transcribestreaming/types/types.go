// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// A list of possible transcriptions for the audio.
type Alternative struct {
	// One or more alternative interpretations of the input audio.
	Items []*Item
	// The text that was transcribed from the audio.
	Transcript *string
}

// Provides a wrapper for the audio chunks that you are sending.
type AudioEvent struct {
	// An audio blob that contains the next part of the audio that you want to
	// transcribe.
	AudioChunk []byte
}

// Represents the audio stream from your application to Amazon Transcribe.
type AudioStream interface {
	isAudioStream()
}

// A blob of audio from your application. You audio stream consists of one or more
// audio events.
type AudioStreamMemberAudioEvent struct {
	Value *AudioEvent
}

func (*AudioStreamMemberAudioEvent) isAudioStream() {}

// A word or phrase transcribed from the input audio.
type Item struct {
	// The type of the item. PRONUNCIATION indicates that the item is a word that was
	// recognized in the input audio. PUNCTUATION indicates that the item was
	// interpreted as a pause in the input audio.
	Type ItemType
	// The offset from the beginning of the audio stream to the end of the audio that
	// resulted in the item.
	EndTime *float64
	// The word or punctuation that was recognized in the input audio.
	Content *string
	// The offset from the beginning of the audio stream to the beginning of the audio
	// that resulted in the item.
	StartTime *float64
	// Indicates whether a word in the item matches a word in the vocabulary filter
	// you've chosen for your real-time stream. If true then a word in the item matches
	// your vocabulary filter.
	VocabularyFilterMatch *bool
}

// The result of transcribing a portion of the input audio stream.
type Result struct {
	// Amazon Transcribe divides the incoming audio stream into segments at natural
	// points in the audio. Transcription results are returned based on these segments.
	// The IsPartial field is true to indicate that Amazon Transcribe has additional
	// transcription data to send, false to indicate that this is the last
	// transcription result for the segment.
	IsPartial *bool
	// The offset in seconds from the beginning of the audio stream to the end of the
	// result.
	EndTime *float64
	// The offset in seconds from the beginning of the audio stream to the beginning of
	// the result.
	StartTime *float64
	// A list of possible transcriptions for the audio. Each alternative typically
	// contains one item that contains the result of the transcription.
	Alternatives []*Alternative
	// A unique identifier for the result.
	ResultId *string
}

// The transcription in a TranscriptEvent ().
type Transcript struct {
	// Result () objects that contain the results of transcribing a portion of the
	// input audio stream. The array can be empty.
	Results []*Result
}

// Represents a set of transcription results from the server to the client. It
// contains one or more segments of the transcription.
type TranscriptEvent struct {
	// The transcription of the audio stream. The transcription is composed of all of
	// the items in the results list.
	Transcript *Transcript
}

// Represents the transcription result stream from Amazon Transcribe to your
// application.
type TranscriptResultStream interface {
	isTranscriptResultStream()
}

// A portion of the transcription of the audio stream. Events are sent periodically
// from Amazon Transcribe to your application. The event can be a partial
// transcription of a section of the audio stream, or it can be the entire
// transcription of that portion of the audio stream.
type TranscriptResultStreamMemberTranscriptEvent struct {
	Value *TranscriptEvent
}

func (*TranscriptResultStreamMemberTranscriptEvent) isTranscriptResultStream() {}

// A client error occurred when the stream was created. Check the parameters of the
// request and try your request again.
type TranscriptResultStreamMemberBadRequestException struct {
	Value *BadRequestException
}

func (*TranscriptResultStreamMemberBadRequestException) isTranscriptResultStream() {}

// A problem occurred while processing the audio. Amazon Transcribe terminated
// processing.
type TranscriptResultStreamMemberInternalFailureException struct {
	Value *InternalFailureException
}

func (*TranscriptResultStreamMemberInternalFailureException) isTranscriptResultStream() {}

// A new stream started with the same session ID. The current stream has been
// terminated.
type TranscriptResultStreamMemberConflictException struct {
	Value *ConflictException
}

func (*TranscriptResultStreamMemberConflictException) isTranscriptResultStream() {}

// Service is currently unavailable. Try your request later.
type TranscriptResultStreamMemberServiceUnavailableException struct {
	Value *ServiceUnavailableException
}

func (*TranscriptResultStreamMemberServiceUnavailableException) isTranscriptResultStream() {}

// Your client has exceeded one of the Amazon Transcribe limits, typically the
// limit on audio length. Break your audio stream into smaller chunks and try your
// request again.
type TranscriptResultStreamMemberLimitExceededException struct {
	Value *LimitExceededException
}

func (*TranscriptResultStreamMemberLimitExceededException) isTranscriptResultStream() {}

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte
}

func (*UnknownUnionMember) isAudioStream()            {}
func (*UnknownUnionMember) isTranscriptResultStream() {}