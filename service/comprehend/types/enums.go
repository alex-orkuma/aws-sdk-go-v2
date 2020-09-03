// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DocumentClassifierMode string

// Enum values for DocumentClassifierMode
const (
	DocumentClassifierModeMulti_class DocumentClassifierMode = "MULTI_CLASS"
	DocumentClassifierModeMulti_label DocumentClassifierMode = "MULTI_LABEL"
)

type EndpointStatus string

// Enum values for EndpointStatus
const (
	EndpointStatusCreating   EndpointStatus = "CREATING"
	EndpointStatusDeleting   EndpointStatus = "DELETING"
	EndpointStatusFailed     EndpointStatus = "FAILED"
	EndpointStatusIn_service EndpointStatus = "IN_SERVICE"
	EndpointStatusUpdating   EndpointStatus = "UPDATING"
)

type EntityType string

// Enum values for EntityType
const (
	EntityTypePerson          EntityType = "PERSON"
	EntityTypeLocation        EntityType = "LOCATION"
	EntityTypeOrganization    EntityType = "ORGANIZATION"
	EntityTypeCommercial_item EntityType = "COMMERCIAL_ITEM"
	EntityTypeEvent           EntityType = "EVENT"
	EntityTypeDate            EntityType = "DATE"
	EntityTypeQuantity        EntityType = "QUANTITY"
	EntityTypeTitle           EntityType = "TITLE"
	EntityTypeOther           EntityType = "OTHER"
)

type InputFormat string

// Enum values for InputFormat
const (
	InputFormatOne_doc_per_file InputFormat = "ONE_DOC_PER_FILE"
	InputFormatOne_doc_per_line InputFormat = "ONE_DOC_PER_LINE"
)

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusSubmitted      JobStatus = "SUBMITTED"
	JobStatusIn_progress    JobStatus = "IN_PROGRESS"
	JobStatusCompleted      JobStatus = "COMPLETED"
	JobStatusFailed         JobStatus = "FAILED"
	JobStatusStop_requested JobStatus = "STOP_REQUESTED"
	JobStatusStopped        JobStatus = "STOPPED"
)

type LanguageCode string

// Enum values for LanguageCode
const (
	LanguageCodeEn    LanguageCode = "en"
	LanguageCodeEs    LanguageCode = "es"
	LanguageCodeFr    LanguageCode = "fr"
	LanguageCodeDe    LanguageCode = "de"
	LanguageCodeIt    LanguageCode = "it"
	LanguageCodePt    LanguageCode = "pt"
	LanguageCodeAr    LanguageCode = "ar"
	LanguageCodeHi    LanguageCode = "hi"
	LanguageCodeJa    LanguageCode = "ja"
	LanguageCodeKo    LanguageCode = "ko"
	LanguageCodeZh    LanguageCode = "zh"
	LanguageCodeZh_tw LanguageCode = "zh-TW"
)

type ModelStatus string

// Enum values for ModelStatus
const (
	ModelStatusSubmitted      ModelStatus = "SUBMITTED"
	ModelStatusTraining       ModelStatus = "TRAINING"
	ModelStatusDeleting       ModelStatus = "DELETING"
	ModelStatusStop_requested ModelStatus = "STOP_REQUESTED"
	ModelStatusStopped        ModelStatus = "STOPPED"
	ModelStatusIn_error       ModelStatus = "IN_ERROR"
	ModelStatusTrained        ModelStatus = "TRAINED"
)

type PartOfSpeechTagType string

// Enum values for PartOfSpeechTagType
const (
	PartOfSpeechTagTypeAdj   PartOfSpeechTagType = "ADJ"
	PartOfSpeechTagTypeAdp   PartOfSpeechTagType = "ADP"
	PartOfSpeechTagTypeAdv   PartOfSpeechTagType = "ADV"
	PartOfSpeechTagTypeAux   PartOfSpeechTagType = "AUX"
	PartOfSpeechTagTypeConj  PartOfSpeechTagType = "CONJ"
	PartOfSpeechTagTypeCconj PartOfSpeechTagType = "CCONJ"
	PartOfSpeechTagTypeDet   PartOfSpeechTagType = "DET"
	PartOfSpeechTagTypeIntj  PartOfSpeechTagType = "INTJ"
	PartOfSpeechTagTypeNoun  PartOfSpeechTagType = "NOUN"
	PartOfSpeechTagTypeNum   PartOfSpeechTagType = "NUM"
	PartOfSpeechTagTypeO     PartOfSpeechTagType = "O"
	PartOfSpeechTagTypePart  PartOfSpeechTagType = "PART"
	PartOfSpeechTagTypePron  PartOfSpeechTagType = "PRON"
	PartOfSpeechTagTypePropn PartOfSpeechTagType = "PROPN"
	PartOfSpeechTagTypePunct PartOfSpeechTagType = "PUNCT"
	PartOfSpeechTagTypeSconj PartOfSpeechTagType = "SCONJ"
	PartOfSpeechTagTypeSym   PartOfSpeechTagType = "SYM"
	PartOfSpeechTagTypeVerb  PartOfSpeechTagType = "VERB"
)

type SentimentType string

// Enum values for SentimentType
const (
	SentimentTypePositive SentimentType = "POSITIVE"
	SentimentTypeNegative SentimentType = "NEGATIVE"
	SentimentTypeNeutral  SentimentType = "NEUTRAL"
	SentimentTypeMixed    SentimentType = "MIXED"
)

type SyntaxLanguageCode string

// Enum values for SyntaxLanguageCode
const (
	SyntaxLanguageCodeEn SyntaxLanguageCode = "en"
	SyntaxLanguageCodeEs SyntaxLanguageCode = "es"
	SyntaxLanguageCodeFr SyntaxLanguageCode = "fr"
	SyntaxLanguageCodeDe SyntaxLanguageCode = "de"
	SyntaxLanguageCodeIt SyntaxLanguageCode = "it"
	SyntaxLanguageCodePt SyntaxLanguageCode = "pt"
)
