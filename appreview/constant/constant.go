package constant

const (
	AppStoreURL  = "https://apps.apple.com/app/id%s"                  // 替换为你的应用 App Store URL
	PlayStoreURL = "https://play.google.com/store/apps/details?id=%s" // 替换为你的应用 Play Store URL
	Android      = "android"
	Ios          = "ios"
)

// 可添加枚举类型增强可读性（示例）
type ReviewStatus int

const (
	ReviewError    ReviewStatus = -1
	ReviewPending  ReviewStatus = 0
	ReviewApproved ReviewStatus = 1
	ReviewRejected ReviewStatus = 2
	ReviewExpired  ReviewStatus = 3
)

// 使用时进行类型转换
// record.Status = int(ReviewApproved)

type TaskStatus int

const (
	TaskError    TaskStatus = -1
	TaskNotStart TaskStatus = 0
	TaskRunning  TaskStatus = 1
	TaskFinish   TaskStatus = 2
	TaskStop     TaskStatus = 3
)

// 使用时进行类型转换
// record.tastStatus = int(Running)
