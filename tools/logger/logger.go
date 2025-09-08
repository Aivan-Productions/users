package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/Aivan-Productions/users/tools/configuration"
)

// Logger обёртка над slog.Logger с дополнительными возможностями
type Logger struct {
	*slog.Logger
	config    *configuration.Config
	transport Transport
}

// Transport интерфейс для отправки логов в различные системы
type Transport interface {
	Send(record slog.Record) error
}

// GrafanaTransport заглушка для будущей реализации Grafana/Loki
type GrafanaTransport struct{}

func (g *GrafanaTransport) Send(record slog.Record) error {
	// TODO: реализовать отправку в Grafana/Loki
	return nil
}

func NewGrafanaTransport() Transport {
	return &GrafanaTransport{}
}

// ColorHandler обработчик с цветным выводом для локального развития
type ColorHandler struct {
	writer io.Writer
	level  slog.Level
}

func NewColorHandler(writer io.Writer, level slog.Level) *ColorHandler {
	return &ColorHandler{
		writer: writer,
		level:  level,
	}
}

func (h *ColorHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *ColorHandler) Handle(ctx context.Context, record slog.Record) error {
	// ANSI color codes
	const (
		reset   = "\033[0m"
		red     = "\033[31m"
		green   = "\033[32m"
		yellow  = "\033[33m"
		blue    = "\033[34m"
		magenta = "\033[35m"
		cyan    = "\033[36m"
	)

	// Выбираем цвет в зависимости от уровня
	var color string
	switch record.Level {
	case slog.LevelDebug:
		color = blue
	case slog.LevelInfo:
		color = green
	case slog.LevelWarn:
		color = yellow
	case slog.LevelError:
		color = red
	default:
		color = reset
	}

	// Форматируем время
	timeStr := record.Time.Format("2006-01-02 15:04:05.000")

	// Форматируем сообщение
	message := fmt.Sprintf("%s%s | %-6s | %s%s\n",
		color,
		timeStr,
		record.Level.String(),
		record.Message,
		reset,
	)

	// Добавляем атрибуты
	if record.NumAttrs() > 0 {
		attrs := make([]string, 0)
		record.Attrs(func(attr slog.Attr) bool {
			attrs = append(attrs, fmt.Sprintf("%s=%v", attr.Key, attr.Value))
			return true
		})
		message += fmt.Sprintf("%s   | %-6s | Attributes: %v%s\n",
			color, "", attrs, reset)
	}

	_, err := h.writer.Write([]byte(message))
	return err
}

func (h *ColorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Для простоты возвращаем тот же handler
	return h
}

func (h *ColorHandler) WithGroup(name string) slog.Handler {
	// Для простоты возвращаем тот же handler
	return h
}

// Init инициализирует логгер на основе конфигурации
func Init(config *configuration.Config) *Logger {
	var handler slog.Handler

	switch config.LOG_FORMAT {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	default: // text
		handler = NewColorHandler(os.Stdout, slog.LevelDebug)
	}

	// Инициализируем транспорт (пока заглушка)
	transport := NewGrafanaTransport()

	return &Logger{
		Logger:    slog.New(handler),
		config:    config,
		transport: transport,
	}
}

// Send отправляет лог через транспорт (для будущего использования)
func (l *Logger) Send(record slog.Record) error {
	return l.transport.Send(record)
}

// With добавляет атрибуты к логгеру
func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		Logger:    l.Logger.With(args...),
		config:    l.config,
		transport: l.transport,
	}
}
