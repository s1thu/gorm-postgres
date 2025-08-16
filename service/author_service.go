package service

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/s1thu/gorm-postgres/model"
	"github.com/s1thu/gorm-postgres/repository"
)

type AuthorService interface {
	CreateAuthor(name, email, bio string) (*model.Author, error)
	GetAuthorByID(id uint) (*model.Author, error)
	GetAllAuthors() ([]*model.Author, error)
	UpdateAuthor(id uint, name, email, bio string) (*model.Author, error)
	DeleteAuthor(id uint) error
	GetAuthorByEmail(email string) (*model.Author, error)
	GetAuthorWithBooks(id uint) (*model.Author, error)
}

type authorServiceImpl struct {
	authorRepo *repository.AuthorRepository
}

func NewAuthorService(authorRepo *repository.AuthorRepository) AuthorService {
	return &authorServiceImpl{
		authorRepo: authorRepo,
	}
}

// CreateAuthor creates a new author with validation
func (s *authorServiceImpl) CreateAuthor(name, email, bio string) (*model.Author, error) {
	// Validate input
	if err := s.validateName(name); err != nil {
		return nil, err
	}

	if err := s.validateEmail(email); err != nil {
		return nil, err
	}

	// Sanitize input
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(strings.ToLower(email))
	bio = strings.TrimSpace(bio)

	// Create new author
	author := &model.Author{
		Name:  name,
		Email: email,
		Bio:   bio,
	}

	err := s.authorRepo.Base.Create(author)
	if err != nil {
		return nil, fmt.Errorf("failed to create author: %w", err)
	}

	return author, nil
}

// GetAuthorByID retrieves an author by their ID
func (s *authorServiceImpl) GetAuthorByID(id uint) (*model.Author, error) {
	if id == 0 {
		return nil, errors.New("invalid author ID")
	}

	author, err := s.authorRepo.Base.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("author with ID %d not found", id)
	}

	return author, nil
}

// GetAllAuthors retrieves all authors
func (s *authorServiceImpl) GetAllAuthors() ([]*model.Author, error) {
	authors, err := s.authorRepo.Base.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve authors: %w", err)
	}

	return authors, nil
}

// UpdateAuthor updates an existing author
func (s *authorServiceImpl) UpdateAuthor(id uint, name, email, bio string) (*model.Author, error) {
	// Validate ID
	if id == 0 {
		return nil, errors.New("invalid author ID")
	}

	// Validate input
	if err := s.validateName(name); err != nil {
		return nil, err
	}

	if err := s.validateEmail(email); err != nil {
		return nil, err
	}

	// Check if author exists
	existingAuthor, err := s.authorRepo.Base.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("author with ID %d not found", id)
	}

	// Sanitize input
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(strings.ToLower(email))
	bio = strings.TrimSpace(bio)

	// Update author fields
	existingAuthor.Name = name
	existingAuthor.Email = email
	existingAuthor.Bio = bio

	err = s.authorRepo.Base.Update(existingAuthor)
	if err != nil {
		return nil, fmt.Errorf("failed to update author: %w", err)
	}

	return existingAuthor, nil
}

// DeleteAuthor deletes an author by ID
func (s *authorServiceImpl) DeleteAuthor(id uint) error {
	if id == 0 {
		return errors.New("invalid author ID")
	}

	// Check if author exists
	_, err := s.authorRepo.Base.GetByID(id)
	if err != nil {
		return fmt.Errorf("author with ID %d not found", id)
	}

	err = s.authorRepo.Base.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete author: %w", err)
	}

	return nil
}

// GetAuthorByEmail retrieves an author by their email
func (s *authorServiceImpl) GetAuthorByEmail(email string) (*model.Author, error) {
	if err := s.validateEmail(email); err != nil {
		return nil, err
	}

	email = strings.TrimSpace(strings.ToLower(email))

	// Note: You'll need to add GetByEmail method to your repository
	// For now, this will cause a compile error until you add it
	author, err := s.authorRepo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("author with email %s not found", email)
	}

	return author, nil
}

// GetAuthorWithBooks retrieves an author along with their books
func (s *authorServiceImpl) GetAuthorWithBooks(id uint) (*model.Author, error) {
	if id == 0 {
		return nil, errors.New("invalid author ID")
	}

	// Note: You'll need to add GetWithBooks method to your repository
	// For now, this will cause a compile error until you add it
	author, err := s.authorRepo.GetWithBooks(id)
	if err != nil {
		return nil, fmt.Errorf("author with ID %d not found", id)
	}

	return author, nil
}

// validateName validates the author name
func (s *authorServiceImpl) validateName(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return errors.New("author name cannot be empty")
	}

	if len(name) < 2 {
		return errors.New("author name must be at least 2 characters long")
	}

	if len(name) > 100 {
		return errors.New("author name must be less than 100 characters")
	}

	return nil
}

// validateEmail validates the email format
func (s *authorServiceImpl) validateEmail(email string) error {
	email = strings.TrimSpace(email)

	if email == "" {
		return errors.New("email cannot be empty")
	}

	// Basic email regex pattern
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	if len(email) > 254 {
		return errors.New("email must be less than 254 characters")
	}

	return nil
}
